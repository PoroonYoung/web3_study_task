// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;
contract BinarySearch{
    function binarySearch(uint256 target)public pure returns(uint256 index){
        uint256[] memory array = new uint256[](10);
        array[0] = 0;
        array[1] = 4;
        array[2] = 7;
        array[3] = 12;
        array[4] = 18;
        array[5] = 25;
        array[6] = 33;
        array[7] = 41;
        array[8] = 56;
        array[9] = 78;
        return  half(array, target, 0, 10);
    }

    function half(uint256[] memory array,uint256 target,uint256 start,uint256 end)private pure returns(uint256 index){
            if (start >= end) {
            return type(uint256).max; 
        }
        index =(start+end)/2;
        uint256 mid = array[index];
        if(mid==target){
            return index;
        }else if (mid<target){
            return half(array,target,index+1,end);
        }else{
            return half(array,target,start,index);
        }
    }
}