pragma solidity ^0.5.2;

contract GatacaID{

    struct Entity {
        string status;
        string pubKey;
    }

    mapping (string => Entity) private registry;
    address private owner;

    modifier onlyOwner {
        require(msg.sender == owner, "invalid user");
        _;
    }

    constructor()
    public {
        owner = msg.sender;
    }

    function getEntity(string calldata _did)
    external view
    returns (string memory status, string memory pubKey) {
        Entity memory e = registry[_did];
        require(bytes(e.status).length != 0, "invalid did");
        require(keccak256(abi.encodePacked((e.status))) != keccak256(abi.encodePacked(("revoked"))), "did revoked");
        require(keccak256(abi.encodePacked((e.status))) != keccak256(abi.encodePacked(("suspended"))), "did suspended");
        return (e.status, e.pubKey);
    }

    function createEntity(string memory _did, string memory _pubkey)
    public onlyOwner{
        Entity memory re = registry[_did];
        require(bytes(re.status).length == 0, "did already in use");
        Entity memory e = Entity ("active", _pubkey);
        registry[_did] = e;
    }

    function revokeEntity(string memory _did)
    public onlyOwner{
        Entity storage e = registry[_did];
        require(bytes(e.status).length != 0, "invalid did");
        e.status = "revoked";
    }

    function suspendEntity(string memory _did)
    public onlyOwner{
        Entity storage e = registry[_did];
        require(bytes(e.status).length != 0, "invalid did");
        require(keccak256(abi.encodePacked((e.status))) != keccak256(abi.encodePacked(("revoked"))), "did revoked");
        e.status = "suspended";
    }

    function activateEntity(string memory _did)
    public onlyOwner{
        Entity storage e = registry[_did];
        require(bytes(e.status).length != 0, "invalid did");
        require(keccak256(abi.encodePacked((e.status))) != keccak256(abi.encodePacked(("revoked"))), "did revoked");
        e.status = "active";
    }

}
