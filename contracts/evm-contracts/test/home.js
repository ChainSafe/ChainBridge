const EmbarkJS = require('Embark/EmbarkJS');
const { HomeContract } = require('./embarkArtifacts/contracts');

let accounts = [];

config({}, (err, _accounts) => {
    accounts = _accounts;
});

contract('DepositContract', () => {
    let HomeInstance;

    before(async function () {
        HomeInstance = await HomeContract.deploy({ 
            arguments: [
                [accounts[0], accounts[1]], // bridge validators
                2, // depoist threshold
                2 // validator threshold
            ]
        }).send();
    });

    it('should set constructor value', async () => {
        let validators = await HomeInstance.methods.Validators().call();
        let depositThreshold = await HomeInstance.methods.voteDepositThreshold().call();
        let validatorThreshold = await HomeInstance.methods.voteValidatorThreshold().call();

        assert.strictEqual(parseInt(validators.length, 10), 2);
    });
});