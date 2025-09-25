// SPDX-License-Identifier: MIT
pragma solidity ^0.8.0;

contract Voting {
    // 候选人 => 票数
    mapping(string => uint64) public voteMap;
    // 判断候选人是否存在
    mapping(string => bool) public isCandidate;
    // 存储所有候选人，用于重置
    string[] public candidateList;

    function vote(string calldata name) external {
        if (!isCandidate[name]) {
            isCandidate[name] = true;
            voteMap[name] = 0;
            candidateList.push(name);
        }

        voteMap[name] += 1;
    }

    function getVotes(string calldata name) external view returns (uint64) {
        return voteMap[name];
    }

    function resetVotes() external {
        for (uint256 i = 0; i < candidateList.length; i++) {
            voteMap[candidateList[i]] = 0;
        }
    }
}
