pragma solidity >=0.7.1 <0.9.0;

contract BasicSyntaxContract {
    uint storedData;
    //melakukan set data
    function set(uint x) public {
        storedData = x;
    }
    //menampilkan data
    function get() public view returns (uint) {
        return storedData;
    }
}