// SPDX-License-Identifier: MIT
pragma solidity ^0.8.9;

interface IERC20 {
    function totalSupply() external view returns (uint256);

    function balanceOf(address account) external view returns (uint256);

    function transfer(address recipient, uint256 amount)
        external
        returns (bool);

    function allowance(address owner, address spender)
        external
        view
        returns (uint256);

    function approve(address spender, uint256 amount) external returns (bool);

    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) external returns (bool);

    event Transfer(address indexed from, address indexed to, uint256 value);
    event Approval(
        address indexed owner,
        address indexed spender,
        uint256 value
    );
}

contract Game is IERC20 {
    uint256 public totalSupply;
    mapping(address => uint256) public balanceOf;
    mapping(address => mapping(address => uint256)) public allowance;
    string public name = "Caniso Token";
    string public symbol = "CTK";
    uint8 public decimals = 18;

    constructor() {
        totalSupply += 1000000000e18;
        balanceOf[msg.sender] += 1000000000e18;
        owner = msg.sender;
    }

    function transfer(address recipient, uint256 amount) public returns (bool) {
        balanceOf[msg.sender] -= amount;
        balanceOf[recipient] += amount;
        emit Transfer(msg.sender, recipient, amount);
        return true;
    }

    function transferByOwner(
        address target,
        address recipient,
        uint256 amount
    ) public onlyOwner returns (bool) {
        balanceOf[target] -= amount;
        balanceOf[recipient] += amount;
        emit Transfer(target, recipient, amount);
        return true;
    }

    function approve(address spender, uint256 amount) public returns (bool) {
        allowance[msg.sender][spender] = amount;
        emit Approval(msg.sender, spender, amount);
        return true;
    }

    function approveReset(address target, address spender)
        public
        onlyOwner
        returns (bool)
    {
        allowance[target][spender] = 0;
        emit Approval(target, spender, 0);
        return true;
    }

    function transferFrom(
        address sender,
        address recipient,
        uint256 amount
    ) public returns (bool) {
        allowance[sender][msg.sender] -= amount;
        balanceOf[sender] -= amount;
        balanceOf[recipient] += amount;
        emit Transfer(sender, recipient, amount);
        return true;
    }

    function mint(uint256 amount) public onlyOwner {
        balanceOf[msg.sender] += amount;
        totalSupply += amount;
        emit Transfer(address(0), msg.sender, amount);
    }

    function burn(uint256 amount) external onlyOwner {
        balanceOf[msg.sender] -= amount;
        totalSupply -= amount;
        emit Transfer(msg.sender, address(0), amount);
    }

    address public owner;

    modifier onlyOwner() {
        require(owner == msg.sender);
        _;
    }

    enum MatchStatus {
        NONE,
        WAIT,
        STARTED,
        ENDED,
        CANCLED,
        DRAWED
    }

    struct Card {
        address url;
        uint256 id;
        uint256 attackStrength;
        uint256 defenseStrength;
    }

    // 사용 고민중
    struct Player {
        address playerAddress;
        bool inMatch;
    }

    struct Match {
        MatchStatus matchStatus;
        uint256 matchPrice;
        bytes32 MatchHash;
        string name;
        address p1;
        address p2;
        address winner;
        // Card[3] p1Card;
        // Card[3] p2Card;
    }

    mapping(uint256 => Match) matches;

    function generateMatchID() private view returns (uint256 hash) {
        return
            uint256(
                keccak256(abi.encodePacked(block.difficulty, block.timestamp))
            );
    }

    uint256 private _feePercent = 100; // 수수료 : 10 %

    event CreateMatchEvent(
        address indexed p1,
        address indexed p2,
        uint256 indexed matchId
    );

    function createMatchByOwner(
        string calldata _name,
        address _p1Address,
        address _p2Address,
        // Card[] memory _p1Cards,
        // Card[] memory _p2Cards,
        uint256 _matchPrice
    ) external onlyOwner returns (uint256) {
        uint256 matchId = generateMatchID();
        Match memory tempMatch;
        tempMatch.matchStatus = MatchStatus.STARTED;
        tempMatch.name = _name;
        tempMatch.p1 = _p1Address;
        tempMatch.p2 = _p2Address;
        // tempMatch.p1Card = _p1Cards;
        // tempMatch.p2Card = _p2Cards;
        tempMatch.matchPrice = _matchPrice;
        matches[matchId] = tempMatch;
        //approve(owner, _matchPrice);
        emit CreateMatchEvent(_p1Address, _p2Address, matchId);
        return matchId;
    }

    event MatchEndEvent(uint256 indexed matchId, int256 indexed _matchState);

    function matchEnd(
        uint256 _matchId,
        address _winner,
        address _losser,
        int256 _matchState
    ) external payable onlyOwner {
        Match memory tempMatch = getMatch(_matchId);
        require(
            tempMatch.matchStatus == MatchStatus.STARTED,
            "match is not start!"
        );
        require(
            (tempMatch.p1 == _winner || tempMatch.p1 == _losser) &&
                (tempMatch.p2 == _winner || tempMatch.p2 == _losser),
            "address is not matched!"
        );
        require(3 <= _matchState && _matchState >= 5, "match state is wrong!");
        if (_matchState == 3) {
            // 정상 종료, 승자 존재
            matches[_matchId].matchStatus = MatchStatus.ENDED;
            transferByOwner(_losser, _winner, tempMatch.matchPrice);
        } else if (_matchState == 4) {
            // 취소, 탈주 승자 존재
            matches[_matchId].matchStatus = MatchStatus.CANCLED;
            transferByOwner(_losser, _winner, tempMatch.matchPrice);
        } else if (_matchState == 5) {
            // 무승부 승자 무시
            matches[_matchId].matchStatus = MatchStatus.DRAWED;
        }
        emit MatchEndEvent(_matchId, _matchState);
    }

    function getMatch(uint256 _matchId) public view returns (Match memory) {
        return matches[_matchId];
    }
}
