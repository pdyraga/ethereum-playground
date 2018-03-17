var EchoGreeter = artifacts.require("EchoGreeter");

module.exports = function(deployer) {
  deployer.deploy(EchoGreeter);
};