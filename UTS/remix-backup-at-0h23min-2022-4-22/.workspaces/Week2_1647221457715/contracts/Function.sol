pragma solidity >=0.7.0<0.9.0;

/*
contract HelloWolrd {

    uint hasil;
    // menggunakan pure untuk mengembalikan satu nilai saja
    function cetakHello() public pure returns(string memory){
        return 'HelloWolrd';
    }
    
    function tambah(uint a, uint b) public {
        //lokal
        uint temp = a + b;

        hasil = temp;
    }

    function getHasil() public view returns(uint){
        return hasil;
    }
}
*/

contract PayableContract {
    uint receivedAmount;
    //function untuk menerima ether
    function receivedEther() payable public {
        // msg.value adalah nilai yang dikirimkan oleh account value akan bernilai wei
        receivedAmount = msg.value;
    }

    function getTotalAmount() public view returns (uint) {
        return receivedAmount;
    }

    function tambah(uint a, uint b) public pure
}