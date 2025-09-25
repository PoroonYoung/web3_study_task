// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract IntegerToRoman {

    function toRoman(uint256 value) external pure returns (string memory) {
        require(value > 0 && value <= 3999, "Value must be 1..3999");

        // 罗马数字及对应数值
        string[13] memory symbols = [
            "M",  "CM", "D",  "CD",
            "C",  "XC", "L",  "XL",
            "X",  "IX", "V",  "IV", "I"
        ];
        uint16[13] memory numbers = [
            1000, 900, 500, 400,
            100,  90,  50,  40,
            10,   9,   5,   4,  1
        ];

        bytes memory result = new bytes(0);
        for (uint i = 0; i < numbers.length; i++) {
            while (value >= numbers[i]) {
                value -= numbers[i];
                result = abi.encodePacked(result, symbols[i]);
            }
        }
        return string(result);
    }
}
