var SimpleGreeter = artifacts.require("SimpleGreeter");

module.exports = function(deployer) {
  deployer.deploy(SimpleGreeter, "om nom nom");
};