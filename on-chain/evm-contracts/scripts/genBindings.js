/**
 * Copyright 2020 ChainSafe Systems
 * SPDX-License-Identifier: LGPL-3.0-only
 */

const fs = require("fs");

const BUILD_PATH = "./bindings/";
const ABI_PATH = BUILD_PATH + "abi/"
const BIN_PATH = BUILD_PATH + "bin/"
// Loop through all the files in the temp directory
fs.readdir("./build/contracts", function (err, files) {
    if (err) {
        console.error("Could not list the directory.", err);
        process.exit(1);
    }
    
    fs.rmdirSync(BUILD_PATH, { recursive: true });
    fs.mkdirSync(BUILD_PATH);
    fs.mkdirSync(ABI_PATH);
    fs.mkdirSync(BIN_PATH);

    files.forEach(function (file, index) {
        const basename = file.split(".")[0];
        const path = './build/contracts/' + file
        let rawdata = fs.readFileSync(path);
        let contract = JSON.parse(rawdata);
        let { abi, bytecode } = contract;
        bytecode = bytecode.substring(2);
        if (abi.length === 0) return;
        fs.writeFileSync(ABI_PATH + basename + ".abi"  , JSON.stringify(abi))
        fs.writeFileSync(BIN_PATH + basename + ".bin", bytecode)
    });
});