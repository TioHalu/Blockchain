pragma solidity >=0.7.0<0.9.0;

contract VariableType {
    //state variable
    uint sum;

    //local  varible
    function tambah(uint num1, uint num2) public {
        //local
        uint temp = num1 + num2;

        sum = temp;
    }

    function  getHasil() public view returns(uint){
        return sum;
    }
}