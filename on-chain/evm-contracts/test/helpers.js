/**
 * Copyright 2020 ChainSafe Systems
 * SPDX-License-Identifier: LGPL-3.0-only
 */

const ethers = require("ethers");

const RelayerActionType = {
    Add: 0,
    Remove: 1
};

const Vote = {
    Yes: 0,
    No: 1,
}

const VoteStatus = {
    Inactive: 0,
    Active: 1,
    Finalized: 2
}

const ThresholdType = {
    Relayer: 0,
    Deposit: 1,
}

// TODO format differently
const dummyData = {
    type: "erc20",
    value: 100,
    to: "0x1",
    from: "0x2"
}

const CreateDepositData = (data = dummyData, depositNonce = 0, originChain = 0) => {
    const hash = ethers.utils.keccak256(ethers.utils.toUtf8Bytes(data));
    return [
        hash,
        depositNonce,
        originChain,
    ];
}

module.exports = {
    RelayerActionType,
    Vote,
    VoteStatus,
    CreateDepositData,
    ThresholdType,
}