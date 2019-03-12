pragma solidity >0.4.99 <0.6.0;
 
contract Energy {
    address payable public supplier;
    address payable public consumer; // also sender in our interpretation
    uint public endTime;
    bytes32 public dataHash; // includes also supplier-generated salt
   
    constructor(address payable _supplier, uint _endTime, bytes32 _dataHash) public payable {
        consumer = msg.sender;
        supplier = _supplier;
        endTime = _endTime;
        dataHash = _dataHash;
    }
    function finishSup(uint256 amount, uint8 v, bytes32 r, bytes32 s) public {
        require(msg.sender == supplier);
        require(ecrecover(keccak256(abi.encodePacked(dataHash, bytes32(amount))), v, r, s) == consumer);
        supplier.transfer(amount);
        selfdestruct(consumer);
    }
    function finishExp() public {
        require(now >= endTime);
        selfdestruct(consumer);
    }
}
