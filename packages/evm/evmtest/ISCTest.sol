// Copyright 2020 IOTA Stiftung
// SPDX-License-Identifier: Apache-2.0

pragma solidity >=0.8.5;

import "@iscmagic/ISC.sol";

contract ISCTest {
    ISCError TestError = isc.registerError("TestError");
    uint64 public constant TokensForGas = 500;

    function getChainID() public view returns (ISCChainID) {
        return isc.getChainID();
    }

    function triggerEvent(string memory s) public {
        isc.triggerEvent(s);
    }

    function triggerEventFail(string memory s) public {
        isc.triggerEvent(s);
        revert();
    }

    event EntropyEvent(bytes32 entropy);

    function emitEntropy() public {
        bytes32 e = isc.getEntropy();
        emit EntropyEvent(e);
    }

    event RequestIDEvent(ISCRequestID reqID);

    function emitRequestID() public {
        ISCRequestID memory reqID = isc.getRequestID();
        emit RequestIDEvent(reqID);
    }

    event SenderAccountEvent(ISCAgentID sender);

    function emitSenderAccount() public {
        ISCAgentID memory sender = isc.getSenderAccount();
        emit SenderAccountEvent(sender);
    }

    function sendBaseTokens(L1Address memory receiver, uint64 baseTokens)
        public
    {
        ISCAllowance memory allowance;
        if (baseTokens == 0) {
            allowance = isc.getAllowanceFrom(msg.sender);
        } else {
            allowance.baseTokens = baseTokens;
        }

        isc.takeAllowedFunds(msg.sender, allowance);

        ISCFungibleTokens memory fungibleTokens;
        require(allowance.baseTokens > TokensForGas);
        fungibleTokens.baseTokens = allowance.baseTokens - TokensForGas;

        ISCDict memory params;

        ISCSendMetadata memory metadata;
        metadata.targetContract = isc.hn("accounts");
        metadata.entrypoint = isc.hn("deposit");
        metadata.params = params;

        ISCSendOptions memory options;

        isc.send(receiver, fungibleTokens, true, metadata, options);
    }

    function revertWithVMError() public view {
        revert VMError(TestError);
    }

    function callInccounter() public {
        ISCDict memory params = ISCDict(new ISCDictItem[](1));
        bytes memory int64Encoded42 = hex"2A00000000000000";
        params.items[0] = ISCDictItem("counter", int64Encoded42);
        ISCAllowance memory allowance;
        isc.call(isc.hn("inccounter"), isc.hn("incCounter"), params, allowance);
    }

    function callSendAsNFT(L1Address memory receiver, NFTID id) public {
        ISCFungibleTokens memory fungibleTokens;
        fungibleTokens.baseTokens = 1074;
        fungibleTokens.tokens = new NativeToken[](0);

        ISCSendMetadata memory metadata;
        metadata.entrypoint = ISCHname.wrap(0x1337);
        metadata.targetContract = ISCHname.wrap(0xd34db33f);

        ISCDict memory optParams = ISCDict(new ISCDictItem[](1));
        bytes memory int64Encoded42 = hex"2A00000000000000";
        optParams.items[0] = ISCDictItem("x", int64Encoded42);
        metadata.params = optParams;

        ISCSendOptions memory options;

        isc.sendAsNFT(receiver, fungibleTokens, id, true, metadata, options);
    }

    function makeISCPanic() public {
        // will produce a panic in ISC
        ISCDict memory params;
        ISCAllowance memory allowance;
        isc.call(
            isc.hn("governance"),
            isc.hn("claimChainOwnership"),
            params,
            allowance
        );
    }

    function moveToAccount(
        ISCAgentID memory targetAgentID,
        ISCAllowance memory allowance
    ) public {
        // moves funds owned by the current contract to the targetAgentID
        ISCDict memory params = ISCDict(new ISCDictItem[](2));
        params.items[0] = ISCDictItem("a", targetAgentID.data);
        bytes memory forceOpenAccount = "\xFF";
        params.items[1] = ISCDictItem("c", forceOpenAccount);
        isc.call(
            isc.hn("accounts"),
            isc.hn("transferAllowanceTo"),
            params,
            allowance
        );
    }

    function sendTo(address payable to, uint256 amount) public payable {
        to.transfer(amount);
        isc.triggerEvent(toString(address(this).balance));
    }

    bytes16 private constant _SYMBOLS = "0123456789abcdef";

    /**
     * @dev Converts a `uint256` to its ASCII `string` decimal representation.
     */
    function toString(uint256 value) internal pure returns (string memory) {
        unchecked {
            uint256 length = log10(value) + 1;
            string memory buffer = new string(length);
            uint256 ptr;
            /// @solidity memory-safe-assembly
            assembly {
                ptr := add(buffer, add(32, length))
            }
            while (true) {
                ptr--;
                /// @solidity memory-safe-assembly
                assembly {
                    mstore8(ptr, byte(mod(value, 10), _SYMBOLS))
                }
                value /= 10;
                if (value == 0) break;
            }
            return buffer;
        }
    }

    /**
     * @dev Return the log in base 10, rounded down, of a positive value.
     * Returns 0 if given 0.
     */
    function log10(uint256 value) internal pure returns (uint256) {
        uint256 result = 0;
        unchecked {
            if (value >= 10**64) {
                value /= 10**64;
                result += 64;
            }
            if (value >= 10**32) {
                value /= 10**32;
                result += 32;
            }
            if (value >= 10**16) {
                value /= 10**16;
                result += 16;
            }
            if (value >= 10**8) {
                value /= 10**8;
                result += 8;
            }
            if (value >= 10**4) {
                value /= 10**4;
                result += 4;
            }
            if (value >= 10**2) {
                value /= 10**2;
                result += 2;
            }
            if (value >= 10**1) {
                result += 1;
            }
        }
        return result;
    }
}
