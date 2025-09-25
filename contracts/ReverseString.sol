// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract ReverseString{

    function reverseStringForAsc( string calldata source) public pure returns (string memory){
        bytes memory src = bytes(source);
        bytes memory targetArray = new bytes(src.length);
        for(uint i=0;i<src.length;i++){
            targetArray[i] = src[src.length-i-1];
        }
        return string(targetArray);
    }

}