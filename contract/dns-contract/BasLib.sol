//SPDX-License-Identifier: UNLICENSED
pragma solidity >= 0.5.0;

library SafeMath {

  function mul(uint256 a, uint256 b) internal pure returns (uint256) {

    if (a == 0) {
      return 0;
    }

    uint256 c = a * b;
    require(c / a == b);

    return c;
  }


  function div(uint256 a, uint256 b) internal pure returns (uint256) {
    require(b > 0); // Solidity only automatically asserts when dividing by 0
    uint256 c = a / b;

    return c;
  }

  function sub(uint256 a, uint256 b) internal pure returns (uint256) {
    require(b <= a);
    uint256 c = a - b;

    return c;
  }

  function add(uint256 a, uint256 b) internal pure returns (uint256) {
    uint256 c = a + b;
    require(c >= a);

    return c;
  }

  function mod(uint256 a, uint256 b) internal pure returns (uint256) {
    require(b != 0);
    return a % b;
  }
}


library BasHash {
    
    function Hash(bytes memory name)
                internal
                pure
                returns (bytes32) {

        return keccak256(abi.encodePacked(name));
    }
}


library BasSet{
    
     struct IndexedAsset {
        mapping(bytes32 => uint256) idx;
        bytes32[] assets;
    }

    function find(BasSet.IndexedAsset storage ia,
                bytes32 hash)
                internal
                view
                returns (bool, uint256){

        uint256 idx = ia.idx[hash];
        if (ia.assets.length > idx) {
            return (ia.assets[idx] == hash, idx);
        } else {
            return (false, 0);
        }
    }

    function append(BasSet.IndexedAsset storage ia,
                    bytes32 hash)
                    internal {

        uint256 idx = ia.idx[hash];
        require(ia.assets.length == 0 || (idx == 0 && ia.assets[0] != hash));

        ia.idx[hash] = uint256(ia.assets.length);
        ia.assets.push(hash);
    }


    function trim(BasSet.IndexedAsset storage ia,
                bytes32 hash) 
                internal {
        (bool found, uint256 idx) = find(ia, hash);
        require(found, "record miss");
        deleteByPosition(ia, idx);
    }

    function trimIfExist(BasSet.IndexedAsset storage ia, 
                        bytes32 hash)
                        internal {
        (bool found, uint256 idx) = find(ia, hash);
        
        if(found){
            deleteByPosition(ia, idx);
        }
    }
    
    function deleteByPosition(BasSet.IndexedAsset storage ia, 
                        uint256 idx)
                        internal {
            bytes32 last = ia.assets[ia.assets.length - 1];
            bytes32 hash = ia.assets[idx];
            //overwrite this position with last element
            ia.assets[idx] = last;
            //alse  change the index record of last element
            ia.idx[last] = idx;
            //delete index record of trim element
            delete ia.idx[hash];
            //pop up the last element if array
            ia.assets.pop();                
    }
    
    function counts(BasSet.IndexedAsset storage ia)
                    internal
                    view
                    returns(uint256){
            return  ia.assets.length;                   
    }
    
    /*
    if end set to 0, we will return all records
    else return [start, end) records
    */
    function slice(BasSet.IndexedAsset storage ia,
                    uint256 start,
                    uint256 end)
                    internal
                    view 
                    returns(bytes32[] memory){
            if (end == 0){
                return ia.assets;
            }
            if (start >= end){
                start = end;
            }
            bytes32[] memory result = new bytes32[](end - start);
            uint256 i = 0;
            uint256 l = ia.assets.length;
            while ((start + i) < end){
                if ((start + i) < l){
                    result[i] = ia.assets[start + i];
                }else{
                    result[i] = bytes32(0);
                }
                i++;
            }
            return result;
    }
}

