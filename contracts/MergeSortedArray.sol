// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract MergeSortedArray{

    function mergeSortedArray() public pure returns(int8[] memory ) {
        int8[] memory arrayOne = new int8[](7);
        arrayOne[0] = 1;
        arrayOne[1] = 3;
        arrayOne[2] = 4;
        arrayOne[3] = 5;
        arrayOne[4] = 7;
        arrayOne[5] = 8;
        arrayOne[6] = 10;

        int8[] memory arrayTwo = new int8[](4);
        arrayTwo[0] = 2;
        arrayTwo[1] = 6;
        arrayTwo[2] = 7;
        arrayTwo[3] = 9;

        int8[] memory result = new int8[](arrayOne.length+arrayTwo.length);
       
        uint i = 0;
        uint j = 0; 

        while (i<arrayOne.length&&j<arrayTwo.length){
            if(arrayOne[i]<arrayTwo[j]){
                result[i+j] = arrayOne[i];
                i++;
            }else {
                result[i+j] = arrayTwo[j];
                j++;
            }
        }
        while (i<arrayOne.length){
            result[i+j] = arrayOne[i];
            i++;
        }
        while (j<arrayTwo.length){
            result[i+j] = arrayTwo[j];
            j++;
        }

        return result;
    }
}