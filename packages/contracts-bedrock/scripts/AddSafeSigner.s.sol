// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

import { console } from "forge-std/console.sol";
import { Script } from "forge-std/Script.sol";
import { Safe, SignatureDecoder } from "safe-contracts/Safe.sol";
import { Enum as SafeOps } from "safe-contracts/common/Enum.sol";

contract AddSafeSigner is Script, SignatureDecoder {
    Safe redstoneSafe = Safe(payable(0x70FdbCb066eD3621647Ddf61A1f40aaC6058Bc89));
    address newSigner = address(0x7211399b320a0417286897fCeD1ee4ba1C1771d4);

    /// @notice Make a call from the Safe contract to an arbitrary address with arbitrary data
    function _callViaSafe(Safe _safe, address _target, bytes memory _data) internal {
        // This is the signature format used the caller is also the signer.
        bytes memory signature = abi.encodePacked(uint256(uint160(msg.sender)), bytes32(0), uint8(1));
        uint8 v;
        bytes32 r;
        bytes32 s;
        (v, r, s) = signatureSplit(signature, 0);
        console.log(v);
        console.logBytes32(r);
        console.logBytes32(s);

        _safe.execTransaction({
            to: _target,
            value: 0,
            data: _data,
            operation: SafeOps.Operation.Call,
            safeTxGas: 0,
            baseGas: 0,
            gasPrice: 0,
            gasToken: address(0),
            refundReceiver: payable(address(0)),
            signatures: signature
        });
    }

    function run() public {
        console.log(msg.sender);

        vm.startBroadcast(msg.sender);
        _callViaSafe(redstoneSafe, address(redstoneSafe), abi.encodeCall(redstoneSafe.addOwnerWithThreshold, (newSigner, 1)));
        vm.stopBroadcast();

        address[] memory owners = redstoneSafe.getOwners();
        console.log("num owners", owners.length);
        for(uint256 i; i<owners.length; i++) {
            console.log("owner", owners[i]);
        }
        uint256 threshold = redstoneSafe.getThreshold();
        console.log("threshold", threshold);

        require(threshold == 1, "invalid threshold");
        require(owners.length == 2, "invalid number of owners");
        require(owners[0] == newSigner, "invalid new owner");
        require(owners[1] == msg.sender, "invalid previous owner");
    }
}
