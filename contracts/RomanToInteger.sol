// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract RomanToInteger {
    function toInteger(string calldata roman) external pure returns (uint256 value) {
        bytes memory s = bytes(roman);
        require(s.length > 0, "Empty string");

        // 辅助函数：单个字符转数值
        for (uint256 i = 0; i < s.length; i++) {
            uint256 curr = _valueOf(s[i]);
            require(curr > 0, "Invalid Roman char");

            uint256 next = 0;
            if (i + 1 < s.length) {
                next = _valueOf(s[i + 1]);
                require(next > 0, "Invalid Roman char");
            }

            // 如果当前值 < 下一个值，说明是减法组合，如 IV(4)
            if (next > curr) {
                value += (next - curr);
                i++; // 跳过下一个
            } else {
                value += curr;
            }
        }
        // 范围约束（传统罗马数字 1..3999）
        require(value > 0 && value <= 3999, "Out of range");
    }

    // 将单个罗马字符转为对应数值
    function _valueOf(bytes1 c) internal pure returns (uint256) {
        if (c == "I") return 1;
        if (c == "V") return 5;
        if (c == "X") return 10;
        if (c == "L") return 50;
        if (c == "C") return 100;
        if (c == "D") return 500;
        if (c == "M") return 1000;
        return 0; // 非法字符
    }
}