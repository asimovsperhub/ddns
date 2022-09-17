# dnsDao名字注册合约v2.0.1
### 合约目录dnscontractv2
### 合约列表
* DnsName      -- 主合约，管理顶级名字
* DnsSubName   -- 每个顶级名字都可以生成DnsSubName合约
* DnsPayment   -- dnsDao支持的租用名字的支付方式（比如usdt）
* DnsPrice     -- 租用名字的定价
* DaoTax       -- dnsDao在租用二级名字的税收
* DnsAccoutant -- dnsDao的记账，提现模块
* DnsOwner     -- 记录名字和所有者对应关系、配置合约的权限由DnsOwner的权限控制
* DnsConf      -- 名字的映射配置
---
## DnsName合约  
### 主要数据结构  
    struct NameItem{
        bytes entireName;
        bool openToReg;
        address erc721Addr;
        uint256 expireTime;
        uint256 tokenId;
    }
    mapping(bytes32=>NameItem) public nameStore;
    mapping(uint256=>bytes32) public id2Hash; 
- nameStore: 保存了名字hash对应的名字   
    - NameItem.entireName 是名字的明文
    - NameItem.openToReg 对外开放注册
    - NameItem.erc721Addr 顶级域名开启注册后，生成的ERC721合约地址
    - NameItem.expireTime 名字租期过期时间
    - NameItem.tokenId 顶级名字对应的tokenId
- id2Hash: 保存了TokenId到名字Hash的映射  

### 方法
* MintName方法  
  功能： Mint一个顶级名字    
``` function MintName(address erc20Addr_, string memory entireName_,uint8 year_) external payable ```  
    - 参数erc20Addr_表示使用的支付方式，如果是eth，填全0地址，否则使用erc20地址
    - 参数entireName是当前要mint的名字，比如com。
    - 参数year_表示要租用的时间，以年为单位，合约内部以365天为一年。  
* ChargeName方法  
    功能： 续租顶级名字（过期后的名字也通过这个方法来续租）  
``` function ChargeName(address erc20Addr_, bytes32 nameHash_,uint8 year_, bool transfer_) external payable ``` 
    - 参数erc20Addr_同MintName
    - 参数nameHash_是顶级名字的hash值
    - 参数year_同MintName
    - 参数transfer_，在名字租期过期的情况下，如果为true，名字将转移到续租人的名下，如果为false则不发生转移
* setRegistered 方法  
    功能： 开启二级名字的注册  
``` function setRegistered(bool open_,bytes32 hash_) external ```  
    - 参数open_ 是否开启二级名字注册
    - 参数hash_ 是顶级名字的hash值

* setContract 方法  
    功能：合约部署初始化后，配置DnsName合约需要引用的DnsPrice、DnsAccountant、DaoTax合约地址  
    ``` function setContract(address taxC_,address accountantC_, address priceC_) external onlyOwner ```  
    - 参数taxC_ DaoTax 合约地址  
    - 参数accountantC_ DnsAccountant合约地址  
    - 参数priceC_ DnsPrice 合约地址  
* setBaseUri 方法  
    功能： 设置ERC721 MetaData地址  
    ``` function setBaseUri(string memory baseUri_) external onlyOwner ```  
    - 参数baseUri是一个可访问的web2地址  
* transferOwnership 方法   
    功能： 将合约转移给新的用户地址  
    ``` function transferOwnership(address payable newOwner_) external override onlyOwner ```  
    - 参数newOwner_ 表示新的用户
* transferFrom 方法  
    功能： 将持有的顶级名字转移给另一个地址  
    ``` function transferFrom(address from, address to, uint256 tokenId) public override ```  
    - 参数from 是转移顶级名字的发起方  
    - 参数to 是转移顶级名字的接收方  
    - 参数tokenId 是顶级名字对应的ID  
* safeTransferFrom 方法  
    功能： 将持有的顶级名字转移给另一个地址
    ``` function safeTransferFrom(address from, address to, uint256 tokenId) public override```  
    - 参数from 是转移顶级名字的发起方  
    - 参数to 是转移顶级名字的接收方  
    - 参数tokenId 是顶级名字对应的ID  
* ownerOf 方法  
    功能：获取tokenId对应的持有者地址  
    ```function ownerOf(uint256 tokenId) public view override returns (address)```
    - 参数tokenId，erc721合约分配的ID
* balanceOf 方法  
    功能：获取一个用户下有多少个tokenId  
    ```function balanceOf(address owner) public view override returns (uint256)```  
    - 参数owner是指tokenId拥有者地址
---
## DnsSubName合约  
### 主要数据结构  
    bytes32 public fatherHash;
    address public fatherAddr;
    struct NameItem{
        bytes entireName;
        mapping(bytes32=>bytes) subName;
        uint256 expireTime;
        uint256 tokenId;
    }

    mapping(bytes32=>NameItem) public nameStore;
    mapping(uint256=>bytes32) public id2Hash;
- fatherHash: 本合约是从顶级名字hash为fatherHash的域名生成的
- fatherAddr：生成本合约的父合约地址
- nameStore: 保存了名字hash对应的名字   
    - NameItem.entireName 是二级名字的明文
    - NameItem.expireTime 名字租期过期时间
    - NameItem.tokenId 顶级名字对应的tokenId  
    - NameItem.subName 三级以及三级以上名字（三级名字附属于二级域名）
 - id2Hash： 保存了TokenId到名字Hash的映射 

### 方法
* MintName方法  
  功能： Mint一个顶级名字    
``` function MintName(address erc20Addr_, string memory entireName_,uint8 year_) external payable ```  
    - 参数erc20Addr_表示使用的支付方式，如果是eth，填全0地址，否则使用erc20地址
    - 参数entireName是当前要mint的名字，比如123.com。
    - 参数year_表示要租用的时间，以年为单位，合约内部以365天为一年。  
* ChargeName方法  
    功能： 续租二级名字（过期后的名字也通过这个方法来续租）  
``` function ChargeName(address erc20Addr_, bytes32 nameHash_,uint8 year_, bool transfer_) external payable ``` 
    - 参数erc20Addr_同MintName
    - 参数nameHash_是二级名字的hash值
    - 参数year_同MintName
    - 参数transfer_，在名字租期过期的情况下，如果为true，名字将转移到续租人的名下，如果为false则不发生转移  
* AddSubName方法  
    功能： 增加三级及以上名字 
    ```function AddSubName(bytes32 hash_, string memory entireName_) external```  
    - 参数hash_是二级域名的Hash地址，比如注册（a.b.123.com),那么hash_=hash(123.com)  
    - 参数entireName是三级及以上名字的明文  
* DelSubName方法  
    功能：删除一个三级及以上名字  
    ``` function DelSubName(bytes32 hash_,string memory entireName_) external```    
    - 参数hash_是二级域名的Hash地址，比如注册（a.b.123.com),那么hash_=hash(123.com)  
    - 参数entireName是三级及以上名字的明文 
* setBaseUri 方法  
    功能： 设置ERC721 MetaData地址  
    ``` function setBaseUri(string memory baseUri_) external onlyOwner ```  
    - 参数baseUri是一个可访问的web2地址  
* transferFrom 方法  
    功能： 将持有的名字转移给另一个地址  
    ``` function transferFrom(address from, address to, uint256 tokenId) public override ```  
    - 参数from 是转移名字的发起方  
    - 参数to 是转移名字的接收方  
    - 参数tokenId 是名字对应的ID  
* safeTransferFrom 方法  
    功能： 将持有的名字转移给另一个地址
    ``` function safeTransferFrom(address from, address to, uint256 tokenId) public override```  
    - 参数from 是转移名字的发起方  
    - 参数to 是转移名字的接收方  
    - 参数tokenId 是名字对应的ID  
* ownerOf 方法  
    功能：获取tokenId对应的持有者地址  
    ```function ownerOf(uint256 tokenId) public view override returns (address)```
    - 参数tokenId，erc721合约分配的ID
* balanceOf 方法  
    功能：获取一个用户下有多少个tokenId  
    ```function balanceOf(address owner) public view override returns (uint256)```  
    - 参数owner是指tokenId拥有者地址    
---  
## DnsPayment合约  
DnsPayment合约保存官方提供的支付方式，只有本合约保存的支付方式，才可以用于支付购买名字和续租名字的费用，DnsPrice和DnsTax的税收依赖于本支付方式；
DnsPayment继承自DaoOwner，DaoOwner负责本合约的权限控制。
### 主要数据结构  
    using LibDnsPayment for LibDnsPayment.DnsPaymentList;
    LibDnsPayment.DnsPaymentList private paymentList;

- paymentList: 由lib库提供的保存dnsDao的支付方式。

### 方法
* getPaymentCount方法  
功能：获取配置的支付方式个数  
```function getPaymentCount() external override view returns (uint256)```
* getPayment方法  
功能：获取详细支付方式配置  
```function getPayment(uint256 index_) external override view returns (string memory,string memory,address,uint8,bool)```
    - 参数index_，支付方式在合约中以array方式存放，在已知支付方式个数的情况下，使用下标值来获取详细的支付方式  
    - 最后一个返回值类型为bool，表示是否disable这个支付方式，如果为true表示禁止使用，false表示本支付方式可以正常使用  
* getPaymentByAddr 方法   
功能：通过erc20地址来获取支付方式  
```function getPaymentByAddr(address erc20Addr_) external override view returns(string memory,string memory,uint8,bool)```  
    - 参数erc20Addr_是指支付方式的erc20地址，如果是eth，则直接天address(0)  
    - 最后一个返回值类型为bool，表示是否disable这个支付方式，如果为true表示禁止使用，false表示本支付方式可以正常使用  
* addPayment方法  
功能： 添加一个支付方式,只有dnsDao才能操作      
```function addPayment(address erc20Addr_, uint8 decimal_,string calldata name_, string calldata symbol_) external onlyDaoOwner```    
    - 参数erc20Addr_是指支付方式的erc20地址，如果是eth，则直接天address(0)    
    - 参数decimal_是描述支付方式的最小单位，例如：usdt是6、eth是18  
    - 参数name_是支付方式的名字  
    - 参数symbol_是支付方式的符号  
* disable方法  
功能：禁用支付方式，只有dnsDao才有权限操作  
```function disable(address erc20Addr_) external onlyDaoOwner    ```   
    - 参数erc20Addr_是指支付方式的erc20地址，如果是eth，则直接天address(0)  
* enable方法  
功能：启用支付方式，只有dnsDao才有权限操作  
```function enable(address erc20Addr_) external onlyDaoOwner    ```   
    - 参数erc20Addr_是指支付方式的erc20地址，如果是eth，则直接天address(0)     
---  
## DnsPrice合约  
DnsPrice合约是dnsDao中提供对名字的定价的合约  
DnsPrice继承自DaoOwner，DaoOwner负责本合约的权限控制。
### 主要数据结构  
    using LibDnsPrice for LibDnsPrice.NamePrice;
    mapping(bytes32=>mapping(address=>LibDnsPrice.NamePrice)) public namePrice;
- namePrice: 由lib库提供的保存对名字的规则和定价

### 方法
* AddPrice方法  
功能：将某种支付方式（DnsPayment中允许的方式）添加到价格合约，并给出默认定价。对于dnsDao的定价，hash_值为0；其他顶级名字的定价hash_值为顶级名字的hash值，只有相应的名字持有者才能操作         
```function AddPrice(bytes32 hash_,address erc20Addr_,uint256 defaultPrice_) external onlyNameOwner(hash_)```  
    - 参数hash_是指顶级名字的hash值，对于dnsDao添加的定价，hash_值为全0  
    - 参数erc20Addr_为erc20地址，如果是eth，则填address(0)  
    - 参数defaultPrice_，默认价格  
* AddPriceLen方法  
功能：增加以字符长度为规则的定价，只有响应的名字持有者才能操作      
```function AddPriceLen(bytes32 hash_,address erc20Addr_,uint256 len_, uint256 price_) external onlyNameOwner(hash_)```  
    - 参数hash_是指顶级名字的hash值，对于dnsDao添加的定价，hash_值为全0  
    - 参数erc20Addr_为erc20地址，如果是eth，则填address(0)  
    - 参数len_为名字的字符长度
    - 参数price_，长度对应的价格  
* DelPriceLen方法  
功能：删除以字符长度为规则的定价，只有响应的名字持有者才能操作         
```function DelPriceLen(bytes32 hash_,address erc20Addr_, uint256 len_) external onlyNameOwner(hash_)```  
    - 参数hash_是指顶级名字的hash值，对于dnsDao添加的定价，hash_值为全0  
    - 参数erc20Addr_为erc20地址，如果是eth，则填address(0)  
    - 参数len_为名字的字符长度  
* RemovePrice方法  
功能：删除支付方式定价，只有响应的名字持有者才能操作         
```function RemovePrice(bytes32 hash_,address erc20Addr_) external onlyNameOwner(hash_)```  
    - 参数hash_是指顶级名字的hash值，对于dnsDao添加的定价，hash_值为全0  
    - 参数erc20Addr_为erc20地址，如果是eth，则填address(0)   
* Price方法  
功能：查询指定长度的定价         
```function Price(bytes32 hash_,address erc20Addr_,uint256 len_) external override view returns(uint256)```  
    - 参数hash_是指顶级名字的hash值，对于dnsDao添加的定价，hash_值为全0  
    - 参数erc20Addr_为erc20地址，如果是eth，则填address(0)  
    - 参数len_为名字的字符长度  
---  
## DaoTax合约  
DaoTax合约是dnsDao中提供对二级名字的税收定价的合约，  
DaoTax继承自DaoOwner，DaoOwner负责本合约的权限控制。
### 主要数据结构  
    using LibDaoTax for LibDaoTax.NameTax;
    mapping(address=>LibDaoTax.NameTax) nameTaxStore;
- nameTaxStore: 由lib库提供的保存对名字的规则和税收定价

### 方法
* AddTax方法  
功能：将某种支付方式（DnsPayment中允许的方式）添加到税收合约，并给出默认税收定价。只有dnsDao才有权限操作     
```function AddTax(address erc20Addr_, uint8 taxType_, uint256 defaultTax_, uint256 percent_,uint256 percentBase_) external override onlyDaoOwner```
    - 参数erc20Addr_为erc20地址，如果是eth，则填address(0)  
    - 参数taxType_为0  
    - 参数defaultTax_为默认税收价格  
    - 参数percent_和percentBase_是一种百分比的配置，合约暂时没有使用
* UpdateTax方法 
功能： 参考AddTax方法，此方法用于更新税收默认参数，只有dnsDao才有权限操作     
```function UpdateTax(address erc20Addr_,uint8 taxType_, uint256 defaultTax_, uint256 percent_,uint256 percentBase_) external override onlyDaoOwner    ```
    - 参数同AddTax方法  
* AddTaxLen方法  
功能：根据名字长度添加税收定价。只有dnsDao才有权限操作     
```function AddTaxLen(address erc20Addr_,uint256 len_,uint256 tax_,uint256 percent_) external override onlyDaoOwner```
    - 参数erc20Addr_为erc20地址，如果是eth，则填address(0)  
    - 参数len_表示名字长度  
    - 参数tax_二级名字税收价格  
    - 参数percent_税收的百分比定价（暂时不用）    
* DelTaxLen方法  
功能：删除某个长度的税收定价。只有dnsDao才有权限操作     
```function DelTaxLen(address erc20Addr_,uint256 len_) external override onlyDaoOwner```
    - 参数erc20Addr_为erc20地址，如果是eth，则填address(0)  
    - 参数len_表示名字长度  
 * tax方法  
功能：查询指定长度的税收定价         
```function tax(address erc20Addr_,uint256 len_) external override view returns(uint8,uint256,uint256,uint256)```  
    - 参数erc20Addr_为erc20地址，如果是eth，则填address(0)  
    - 参数len_为名字的字符长度     
---  
## DnsAccountant合约  
DnsAccountant合约为dnsDao和各个顶级名字提供记账能力，  
DnsAccountant继承自DaoOwner，DaoOwner负责本合约的权限控制。
### 主要数据结构  
    using LibDnsAccountant for LibDnsAccountant.NameAccountant;
    LibDnsAccountant.NameAccountant private nameAccountant;
- nameAccountant: 由lib库提供记录账单的数据结构

### 方法  
* withdrawCash方法   
功能：提现  
```function withdrawCash(address erc20Addr_,uint256 amount_,address payable out_, address contractAddr_) external override onlyContractOwner(contractAddr_)```  
    - 参数erc20Addr_为erc20地址，如果是eth，则填address(0)   
    - 参数amount_提现数额  
    - 参数out_提现到哪个账号
    - 参数contractAddr_校验用户是否为本合约地址的owner，如果不是，则不能提取。
* get方法   
功能：查询可提现最大数量  
```function get(address account_,address erc20Addr_) external override view returns(uint256)```  
    - 参数account_,记账地址，顶级名字的erc721地址，对于dnsDao，改地址就是DnsName地址
    - 参数erc20Addr_为erc20地址，如果是eth，则填address(0)        

---  
## DnsConf合约  
DnsConf为每个名字提供映射配置的合约  
DnsConf继承自DaoOwner，DaoOwner负责本合约的权限控制。
### 主要数据结构  
    using LibDnsConf for LibDnsConf.NameConfMap;
    LibDnsConf.NameConfMap private nameConfStore;
- nameConfStore: 由lib库提供映射配置
### 方法
* addType方法   
功能：添加一种映射内容，只有dnsDao才有权限添加  
```function addType(bytes memory tName_) external onlyDaoOwner```    
    - 参数tName_是指映射内容的名字  
* addMap方法   
功能：添加一项映射，只有名字持有者才能添加  
```function addMap( bytes32 nameHash_,bytes memory tName_, bytes memory v_) external onlyNameOwner(nameHash_)```   
    - 参数nameHash_表示名字的hash值
    - 参数tName_是指映射内容的名字  
    - 参数v_值映射内容的值
* removeMap方法    
功能： 删除一项映射，只有名字持有者才能删除  
```function removeMap(bytes32 nameHash_,bytes memory tName_) external onlyNameOwner(nameHash_)```  
    - 参数nameHash_表示名字的hash值
    - 参数tName_是指映射内容的名字  
* updateMap方法   
功能：更新一项映射，只有名字持有者才能添加  
```function updateMap( bytes32 nameHash_,bytes memory tName_, bytes memory v_) external onlyNameOwner(nameHash_)```   
    - 参数nameHash_表示名字的hash值
    - 参数tName_是指映射内容的名字  
    - 参数v_值映射内容的值   


















