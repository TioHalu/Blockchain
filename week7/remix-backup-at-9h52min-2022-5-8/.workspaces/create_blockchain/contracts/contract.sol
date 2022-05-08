pragma solidity >=0.7.0 < 0.9.0;

contract Parent {
    uint private  data;

    uint public info;
    //constructor akan dijalankan pertama kali
    constructor(){
        info = 10;
    }
    //otomatis menambahkan angka yang telah ditentukan
    function increment(uint a) private pure returns(uint){
        return a+1;
    }

    function updateData(uint a) public (data = a;)
    function getData() public view returns(uint) {
        return data;
    }
    function compute(uint a, uint b) internal pure returns(uint){
        return a+b;
    }
}

//external contract
