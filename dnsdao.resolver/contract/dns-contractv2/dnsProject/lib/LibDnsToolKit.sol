//SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

library LibDnsToolKit{
    uint8 constant MAX_NAME_SEG_LEN   = 64;
    uint16 constant MAX_SUB_NAME_LEN = 192;

    bytes1 constant DOT_CHAR_VAL = 0x2e;

    function fatherNameHash(bytes memory name_) public pure returns(bytes32){
        require(!verifyRoot(name_),"no a sub name");

        string memory fName = getFatherName(name_);

        return keccak256(abi.encodePacked(fName));
    }

    function entireNameHash(string memory name_) public pure returns(bytes32){
        return keccak256(abi.encodePacked(name_));
    }

    function getFatherName(bytes memory name_) public pure returns(string memory){
        require(verifyName(name_),"name not valid");
        bytes memory ret;
        uint256 idx;
        for (uint256 i = 0; i < name_.length; i++){
            if(idx>0){
                ret[i-idx] = name_[i];
            }
            if (idx==0 && name_[i] == DOT_CHAR_VAL){
                ret = new bytes(name_.length-i-1);
                idx = i+1;
            }
        }
        require(ret.length>0,"no father name");
        return string(ret);
    }

    function getLevel2FatherName(bytes memory name_) public pure returns(bytes memory){
        require(verifyName(name_),"name not valid");
        uint8 dotCnt;
        uint256 idx;

        for(uint256 i=name_.length-1;i>=0;){
            if (name_[i] == DOT_CHAR_VAL){
                dotCnt ++;
            }
            if(dotCnt>=2){
                idx = i;
                break;
            }
            if(i==0){
                break;
            }
            i--;
        }
        require(dotCnt>=2,"not a level 2 father name");
        bytes memory ret = new bytes(name_.length-idx-1);
        for(uint256 i=idx+1;i<name_.length;i++){
            ret[i-idx-1] = name_[i];
        }
        return ret;
    }

    function getLeve2FatherNameHash(bytes memory name_) public pure returns(bytes32){
        bytes memory l2fatherName = getLevel2FatherName(name_);
        return keccak256(abi.encodePacked(l2fatherName));
    }

    function validChar(bytes1 c)
    internal
    pure
    returns (bool) {
        return
        (c >= 0x30 && c <= 0x39) ||
        (c >= 0x61 && c <= 0x7a) ||
        c == 0x2d ||
        c == 0x5f;
    }

    function getSubName(bytes memory name_) public pure returns(bytes memory){
        bytes memory ret;
        for (uint256 i = 0; i < name_.length; i++){
            if (name_[i] == DOT_CHAR_VAL){
                ret = new bytes(i);
                break;
            }
        }
        for(uint256 i=0;i<ret.length;i++){
            ret[i] = name_[i];
        }
        return ret;
    }


    function verifyName(bytes memory name)
    public
    pure
    returns (bool) {

        if (name.length >= MAX_SUB_NAME_LEN){
            return false;
        }

        if (name[0] == DOT_CHAR_VAL || name[name.length-1] == DOT_CHAR_VAL){
            return false;
        }

        uint256 segementLength = 0;

        for (uint256 i = 0; i < name.length; i++) {

            bytes1 char = name[i];
            if (!validChar(char) && char != DOT_CHAR_VAL){
                return false;
            }

            if (char == DOT_CHAR_VAL){
                if (segementLength == 0){
                    return false;
                }

                segementLength = 0;
                continue;
            }

            segementLength += 1;
            if (segementLength >= MAX_NAME_SEG_LEN){
                return false;
            }
        }

        return true;
    }

    function verifyRoot(bytes memory name)
    public
    pure
    returns (bool) {

        if (name.length == 0 || name.length >= MAX_NAME_SEG_LEN) {
            return false;
        }

        for (uint8 i = 0; i < name.length; i++) {
            if (!validChar(name[i])) {
                return false;
            }
        }

        return true;
    }

    function max(uint256 x,uint256 y) public pure returns(uint256){
        if(x>y){
            return x;
        }
        return y;
    }

    function bytes2Bytes32(bytes memory b) public pure returns (bytes32) {
        bytes32 out;

        for (uint i = 0; i < b.length && i<32; i++) {
            out |= bytes32(b[i] & 0xFF) >> (i * 8);
        }
        return out;
    }

    //    function bytes2Bytes32(bytes memory v_) public pure returns(bytes32){
    //        if(v_.length == 0 || v_.length >32){
    //            return bytes32(0);
    //        }
    //
    //        bytes32 r;
    //
    //        for(uint256 i=0;i<v_.length;i++){
    //            r[i]=v_[i];
    //        }
    //        return r;
    //    }

}