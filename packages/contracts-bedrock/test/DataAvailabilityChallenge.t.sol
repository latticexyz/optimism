// SPDX-License-Identifier: UNLICENSED
pragma solidity 0.8.15;

import { Test, console } from "forge-std/Test.sol";
import { DataAvailabilityChallenge, ChallengeStatus, Challenge } from "../src/L1/DataAvailabilityChallenge.sol";
import { Proxy } from "src/universal/Proxy.sol";

address constant DAC_OWNER = address(1234);
uint256 constant CHALLENGE_WINDOW = 1000;
uint256 constant RESOLVE_WINDOW = 1000;
uint256 constant BOND_SIZE = 1000;

contract DataAvailabilityChallengeTest is Test {
    DataAvailabilityChallenge public dac;

    function setUp() public virtual {
        dac = new DataAvailabilityChallenge();
        dac.initialize(DAC_OWNER, CHALLENGE_WINDOW, RESOLVE_WINDOW, BOND_SIZE);
    }

    function testInitialize() public {
        assertEq(dac.owner(), DAC_OWNER);
        assertEq(dac.challengeWindow(), CHALLENGE_WINDOW);
        assertEq(dac.resolveWindow(), RESOLVE_WINDOW);
        assertEq(dac.bondSize(), BOND_SIZE);

        vm.expectRevert("Initializable: contract is already initialized");
        dac.initialize(DAC_OWNER, CHALLENGE_WINDOW, RESOLVE_WINDOW, BOND_SIZE);
    }

    function testDeposit() public {
        assertEq(dac.balances(address(this)), 0);
        dac.deposit{ value: 1000 }();
        assertEq(dac.balances(address(this)), 1000);
    }

    function testReceive() public {
        assertEq(dac.balances(address(this)), 0);
        (bool success,) = payable(address(dac)).call{ value: 1000 }("");
        assertTrue(success);
        assertEq(dac.balances(address(this)), 1000);
    }

    function testWithdraw(address sender, uint256 amount) public {
        assumePayable(sender);
        assumeNotPrecompile(sender);
        vm.assume(sender != address(dac));
        vm.assume(sender.balance == 0);
        vm.deal(sender, amount);

        vm.prank(sender);
        dac.deposit{ value: amount }();

        assertEq(dac.balances(sender), amount);
        assertEq(sender.balance, 0);

        vm.prank(sender);
        dac.withdraw();

        assertEq(dac.balances(sender), 0);
        assertEq(sender.balance, amount);
    }

    function testChallengeSuccess(address challenger, uint256 challengedBlockNumber, bytes32 challengedHash) public {
        // Assume the challenger is not the 0 address
        vm.assume(challenger != address(0));

        // Assume the block number is not close to the max uint256 value
        vm.assume(challengedBlockNumber < type(uint256).max - dac.challengeWindow() - dac.resolveWindow());
        uint256 requiredBond = dac.bondSize();

        // Move to a block after the challenged block
        vm.roll(challengedBlockNumber + 1);

        // Deposit the required bond
        vm.deal(challenger, requiredBond);
        vm.prank(challenger);
        dac.deposit{ value: requiredBond }();

        // Expect the challenge status to be uninitialized
        assertEq(uint8(dac.getChallengeStatus(challengedBlockNumber, challengedHash)), uint8(ChallengeStatus.Uninitialized));

        // Challenge a (blockNumber,hash) tuple
        vm.prank(challenger);
        dac.challenge(challengedBlockNumber, challengedHash);

        // Challenge should have been created
        (address _challenger, uint256 _lockedBond, uint256 _startBlock, uint256 _resolvedBlock) = dac.challenges(challengedBlockNumber, challengedHash);
        assertEq(_challenger, challenger);
        assertEq(_startBlock, block.number);
        assertEq(_resolvedBlock, 0);
        assertEq(_lockedBond, requiredBond);
        assertEq(uint8(dac.getChallengeStatus(challengedBlockNumber, challengedHash)), uint8(ChallengeStatus.Active));

        // Challenge should have decreased the challenger's bond size
        assertEq(dac.balances(challenger), 0);
    }

    function testChallengeDeposit(address challenger, uint256 challengedBlockNumber, bytes32 challengedHash) public {
        // Assume the challenger is not the 0 address
        vm.assume(challenger != address(0));

        // Assume the block number is not close to the max uint256 value
        vm.assume(challengedBlockNumber < type(uint256).max - dac.challengeWindow() - dac.resolveWindow());
        uint256 requiredBond = dac.bondSize();

        // Move to a block after the challenged block
        vm.roll(challengedBlockNumber + 1);

        // Expect the challenge status to be uninitialized
        assertEq(uint8(dac.getChallengeStatus(challengedBlockNumber, challengedHash)), uint8(ChallengeStatus.Uninitialized));

        // Deposit the required bond as part of the challenge transaction
        vm.deal(challenger, requiredBond);
        vm.prank(challenger);
        dac.challenge{ value: requiredBond }(challengedBlockNumber, challengedHash);

        // Challenge should have been created
        (address _challenger, uint256 _lockedBond, uint256 _startBlock, uint256 _resolvedBlock) = dac.challenges(challengedBlockNumber, challengedHash);
        assertEq(_challenger, challenger);
        assertEq(_startBlock, block.number);
        assertEq(_resolvedBlock, 0);
        assertEq(_lockedBond, requiredBond);
        assertEq(uint8(dac.getChallengeStatus(challengedBlockNumber, challengedHash)), uint8(ChallengeStatus.Active));

        // Challenge should have decreased the challenger's bond size
        assertEq(dac.balances(challenger), 0);
    }

    function testChallengeFailBondTooLow() public {
        uint256 requiredBond = dac.bondSize();
        uint256 actualBond = requiredBond - 1;
        dac.deposit{ value: actualBond }();

        vm.expectRevert(abi.encodeWithSelector(DataAvailabilityChallenge.BondTooLow.selector, actualBond, requiredBond));
        dac.challenge(0, "some hash");
    }

    function testChallengeFailChallengeExists() public {
        // Move to a block after the hash to challenge
        vm.roll(2);

        // First challenge succeeds
        dac.deposit{ value: dac.bondSize() }();
        dac.challenge(0, "some hash");

        // Second challenge of the same hash/blockNumber fails
        dac.deposit{ value: dac.bondSize() }();
        vm.expectRevert(abi.encodeWithSelector(DataAvailabilityChallenge.ChallengeExists.selector));
        dac.challenge(0, "some hash");

        // Challenge succeed if the challenged block number is different
        dac.deposit{ value: dac.bondSize() }();
        dac.challenge(1, "some hash");

        // Challenge succeed if the challenged hash is different
        dac.deposit{ value: dac.bondSize() }();
        dac.challenge(0, "some other hash");
    }

    function testChallengeFailBeforeChallengeWindow() public {
        uint256 challengedBlockNumber = 1;
        bytes32 challengedHash = "some hash";

        // Move to challenged block
        vm.roll(challengedBlockNumber);

        // Challenge fails because the current block number must be after the challenged block
        dac.deposit{ value: dac.bondSize() }();
        vm.expectRevert(abi.encodeWithSelector(DataAvailabilityChallenge.ChallengeWindowNotOpen.selector));
        dac.challenge(challengedBlockNumber, challengedHash);
    }

    function testChallengeFailAfterChallengeWindow() public {
        uint256 challengedBlockNumber = 1;
        bytes32 challengedHash = "some hash";

        // Move to block after the challenge window
        vm.roll(challengedBlockNumber + dac.challengeWindow() + 1);

        // Challenge fails because the block number is after the challenge window
        dac.deposit{ value: dac.bondSize() }();
        vm.expectRevert(abi.encodeWithSelector(DataAvailabilityChallenge.ChallengeWindowNotOpen.selector));
        dac.challenge(challengedBlockNumber, challengedHash);
    }

    function testResolveSuccess(bytes memory preImage, uint256 challengedBlockNumber) public returns (uint256 resolveGasUsed) {
        // Assume the block number is not close to the max uint256 value
        vm.assume(challengedBlockNumber < type(uint256).max - dac.challengeWindow() - dac.resolveWindow());
        bytes32 challengedHash = keccak256(preImage);

        // Move to block after challenged block
        vm.roll(challengedBlockNumber + 1);

        // Challenge the hash
        dac.deposit{ value: dac.bondSize() }();
        dac.challenge(challengedBlockNumber, challengedHash);

        // Resolve the challenge
        uint256 gasLeftBefore = gasleft();
        dac.resolve(challengedBlockNumber, challengedHash, preImage);
        resolveGasUsed = gasLeftBefore - gasleft();

        // Expect the challenge to be resolved
        (address _challenger, uint256 _lockedBond, uint256 _startBlock, uint256 _resolvedBlock) = dac.challenges(challengedBlockNumber, challengedHash);

        assertEq(_challenger, address(this));
        assertEq(_lockedBond, dac.bondSize());
        assertEq(_startBlock, block.number);
        assertEq(_resolvedBlock, block.number);
        assertEq(uint8(dac.getChallengeStatus(challengedBlockNumber, challengedHash)), uint8(ChallengeStatus.Resolved));
    }

    function testResolveFailNonExistentChallenge() public {
        bytes memory preImage = "some preimage";
        uint256 challengedBlockNumber = 1;

        // Move to block after challenged block
        vm.roll(challengedBlockNumber + 1);

        // Resolving a non-existent challenge fails
        vm.expectRevert(abi.encodeWithSelector(DataAvailabilityChallenge.ChallengeNotActive.selector));
        dac.resolve(challengedBlockNumber, keccak256(preImage), preImage);
    }

    function testResolveFailResolved() public {
        bytes memory preImage = "some preimage";
        bytes32 challengedHash = keccak256(preImage);
        uint256 challengedBlockNumber = 1;

        // Move to block after challenged block
        vm.roll(challengedBlockNumber + 1);

        // Challenge the hash
        dac.deposit{ value: dac.bondSize() }();
        dac.challenge(challengedBlockNumber, challengedHash);

        // Resolve the challenge
        dac.resolve(challengedBlockNumber, challengedHash, preImage);

        // Resolving an already resolved challenge fails
        vm.expectRevert(abi.encodeWithSelector(DataAvailabilityChallenge.ChallengeNotActive.selector));
        dac.resolve(challengedBlockNumber, challengedHash, preImage);
    }

    function testResolveFailExpired() public {
        bytes memory preImage = "some preimage";
        bytes32 challengedHash = keccak256(preImage);
        uint256 challengedBlockNumber = 1;

        // Move to block after challenged block
        vm.roll(challengedBlockNumber + 1);

        // Challenge the hash
        dac.deposit{ value: dac.bondSize() }();
        dac.challenge(challengedBlockNumber, challengedHash);

        // Move to a block after the resolve window
        vm.roll(block.number + dac.resolveWindow() + 1);


        // Resolving an expired challenge fails
        vm.expectRevert(abi.encodeWithSelector(DataAvailabilityChallenge.ChallengeNotActive.selector));
        dac.resolve(challengedBlockNumber, challengedHash, preImage);
    }

    function testResolveFailAfterResolveWindow() public {
        bytes memory preImage = "some preimage";
        bytes32 challengedHash = keccak256(preImage);
        uint256 challengedBlockNumber = 1;

        // Move to block after challenged block
        vm.roll(challengedBlockNumber + 1);

        // Challenge the hash
        dac.deposit{ value: dac.bondSize() }();
        dac.challenge(challengedBlockNumber, challengedHash);

        // Move to block after resolve window
        vm.roll(block.number + dac.resolveWindow() + 1);

        // Resolve the challenge
        vm.expectRevert(abi.encodeWithSelector(DataAvailabilityChallenge.ChallengeNotActive.selector));
        dac.resolve(challengedBlockNumber, challengedHash, preImage);
    }

    function testUnlockBondSuccess(bytes memory preImage, uint256 challengedBlockNumber) public {
        // Assume the block number is not close to the max uint256 value
        vm.assume(challengedBlockNumber < type(uint256).max - dac.challengeWindow() - dac.resolveWindow());
        bytes32 challengedHash = keccak256(preImage);

        // Move to block after challenged block
        vm.roll(challengedBlockNumber + 1);

        // Challenge the hash
        dac.deposit{ value: dac.bondSize() }();
        dac.challenge(challengedBlockNumber, challengedHash);

        // Move to a block after the resolve window
        vm.roll(block.number + dac.resolveWindow() + 1);

        uint256 balanceBeforeUnlock = dac.balances(address(this));

        // Unlock the bond associated with the challenge
        dac.unlockBond(challengedBlockNumber, challengedHash);

        // Expect the balance to be increased by the bond size
        uint256 balanceAfterUnlock = dac.balances(address(this));
        assertEq(balanceAfterUnlock, balanceBeforeUnlock + dac.bondSize());

        // Expect the bond to be unlocked
        (address _challenger, uint256 _lockedBond, uint256 _startBlock, uint256 _resolvedBlock) = dac.challenges(challengedBlockNumber, challengedHash);

        assertEq(_challenger, address(this));
        assertEq(_lockedBond, 0);
        assertEq(_startBlock, challengedBlockNumber + 1);
        assertEq(_resolvedBlock, 0);
        assertEq(uint8(dac.getChallengeStatus(challengedBlockNumber, challengedHash)), uint8(ChallengeStatus.Expired));

        // Unlock the bond again, expect the balance to remain the same
        dac.unlockBond(challengedBlockNumber, challengedHash);
        assertEq(dac.balances(address(this)), balanceAfterUnlock);

    }

    function testUnlockBondFailNonExistentChallenge() public {
        bytes memory preImage = "some preimage";
        bytes32 challengedHash = keccak256(preImage);
        uint256 challengedBlockNumber = 1;

        // Move to block after challenged block
        vm.roll(challengedBlockNumber + 1);

        // Unlock a bond of a non-existent challenge fails
        vm.expectRevert(abi.encodeWithSelector(DataAvailabilityChallenge.ChallengeNotExpired.selector));
        dac.unlockBond(challengedBlockNumber, challengedHash);
    }

    function testUnlockBondFailResolvedChallenge() public {
        bytes memory preImage = "some preimage";
        bytes32 challengedHash = keccak256(preImage);
        uint256 challengedBlockNumber = 1;

        // Move to block after challenged block
        vm.roll(challengedBlockNumber + 1);

        // Challenge the hash
        dac.deposit{ value: dac.bondSize() }();
        dac.challenge(challengedBlockNumber, challengedHash);

        // Resolve the challenge
        dac.resolve(challengedBlockNumber, challengedHash, preImage);

        // Attempting to unlock a bond of a resolved challenge fails
        vm.expectRevert(abi.encodeWithSelector(DataAvailabilityChallenge.ChallengeNotExpired.selector));
        dac.unlockBond(challengedBlockNumber, challengedHash);
    }

    function testUnlockBondExpiredChallengeTwice() public {
        bytes memory preImage = "some preimage";
        bytes32 challengedHash = keccak256(preImage);
        uint256 challengedBlockNumber = 1;

        // Move to block after challenged block
        vm.roll(challengedBlockNumber + 1);

        // Challenge the hash
        dac.deposit{ value: dac.bondSize() }();
        dac.challenge(challengedBlockNumber, challengedHash);

        // Move to a block after the challenge window
        vm.roll(block.number + dac.resolveWindow() + 1);

        // Unlock the bond
        dac.unlockBond(challengedBlockNumber, challengedHash);

        uint256 balanceAfterUnlock = dac.balances(address(this));

        // Unlock the bond again doesn't change the balance
        dac.unlockBond(challengedBlockNumber, challengedHash);
        assertEq(dac.balances(address(this)), balanceAfterUnlock);
    }

    function testUnlockFailResolveWindowNotClosed() public {
        bytes memory preImage = "some preimage";
        bytes32 challengedHash = keccak256(preImage);
        uint256 challengedBlockNumber = 1;

        // Move to block after challenged block
        vm.roll(challengedBlockNumber + 1);

        // Challenge the hash
        dac.deposit{ value: dac.bondSize() }();
        dac.challenge(challengedBlockNumber, challengedHash);

        vm.roll(block.number + dac.resolveWindow() - 1);

        // Expiring the challenge before the resolve window closes fails
        vm.expectRevert(abi.encodeWithSelector(DataAvailabilityChallenge.ChallengeNotExpired.selector));
        dac.unlockBond(challengedBlockNumber, challengedHash);
    }

    function testSetBondSize() public {
        uint256 requiredBond = dac.bondSize();
        uint256 actualBond = requiredBond - 1;
        dac.deposit{ value: actualBond }();

        // Expect the challenge to fail because the bond is too low
        vm.expectRevert(abi.encodeWithSelector(DataAvailabilityChallenge.BondTooLow.selector, actualBond, requiredBond));
        dac.challenge(0, "some hash");

        // Reduce the required bond
        vm.prank(DAC_OWNER);
        dac.setBondSize(actualBond);

        // Expect the challenge to succeed
        dac.challenge(0, "some hash");
    }

    function testSetBondSizeFailOnlyOwner(address notOwner, uint256 newBondSize) public {
        vm.assume(notOwner != DAC_OWNER);

        // Expect setting the bond size to fail because the sender is not the owner
        vm.prank(notOwner);
        vm.expectRevert("Ownable: caller is not the owner");
        dac.setBondSize(newBondSize);
    }
}

contract DataAvailabilityChallengeProxyTest is DataAvailabilityChallengeTest {
    function setUp() public virtual override {
        Proxy proxy = new Proxy(address(this));
        proxy.upgradeTo(address(new DataAvailabilityChallenge()));
        dac = DataAvailabilityChallenge(payable(address(proxy)));
        dac.initialize(DAC_OWNER, CHALLENGE_WINDOW, RESOLVE_WINDOW, BOND_SIZE);
    }

    function generateData(uint256 size, bool zeroBytes) internal pure returns (bytes memory data) {
        data = new bytes(size);
        if(zeroBytes) return data;

        for(uint256 i; i < size; i++) {
            data[i] = hex'FF';
        }
    }

    function logGasCost(uint256 size, bool zeroBytes) internal {
        bytes memory data = generateData(size, zeroBytes);
        uint256 gasUsed = testResolveSuccess(data, 1);
        console.log("gas used", size, zeroBytes, gasUsed);
    }

    function testResolveGasZeroBytesLength0() public {
        logGasCost(0, true);
    }

    function testResolveGasNonZeroBytesLength0() public {
        logGasCost(0, false);
    }

    function testResolveGasZeroBytesLength32() public {
        logGasCost(32, true);
    }

    function testResolveGasNonZeroBytesLength32() public {
        logGasCost(32, false);
    }

    function testResolveGasZeroBytesLength1000() public {
        logGasCost(1000, true);
    }

    function testResolveGasNonZeroBytesLength1000() public {
        logGasCost(1000, false);
    }

    function testResolveGasZeroBytesLength10000() public {
        logGasCost(10000, true);
    }

    function testResolveGasNonZeroBytesLength10000() public {
        logGasCost(10000, false);
    }

    function testResolveGasZeroBytesLength128000() public {
        logGasCost(128000, true);
    }

    function testResolveGasNonZeroBytesLength128000() public {
        logGasCost(128000, false);
    }
}
