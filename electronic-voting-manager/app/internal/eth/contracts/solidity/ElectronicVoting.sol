// SPDX-License-Identifier: UNLICENSED
pragma solidity ^0.8.0;

contract ElectronicVoting {

    string  votingTitle;
    uint    numberRegisteredVoters;
    uint    votingEndTime;
    address owner;
    bool    optionsCompleted;

    modifier onlyOwner
    {
        if (msg.sender != owner)
            revert();
        _;
    }

    constructor(string memory _votingTitle, uint _votingEndTime)
    {
        if (block.timestamp > _votingEndTime)
            revert();

        votingTitle             = _votingTitle;
        numberRegisteredVoters  = 0;
        votingEndTime           = _votingEndTime;
        owner                   = msg.sender;
        optionsCompleted        = false;
    }

    struct VotingOption
    {
        string name;
        uint numberVotes;
    }

    VotingOption[] public votingOptions;

    function addVotingOption(string memory _votingOptionName) public onlyOwner
    {
        if (block.timestamp > votingEndTime || optionsCompleted == true)
            revert();

        votingOptions.push(VotingOption({
            name: _votingOptionName,
            numberVotes: 0
        }));
    }

    function completeVotingOptions() public onlyOwner
    {
        if (block.timestamp > votingEndTime || votingOptions.length < 2)
            revert();

        optionsCompleted = true;
    }

    struct Voter
    {
        bool hasRightToVote;
        bool voted;
        uint votedFor;
    }

    mapping(address => Voter) public voters;

    function giveRightToVote(address _voterAddress) public onlyOwner
    {
        if (block.timestamp > votingEndTime)
            revert();
        voters[_voterAddress].hasRightToVote = true;
        numberRegisteredVoters += 1;
    }

    function vote(uint idx) public
    {
        if(block.timestamp > votingEndTime || optionsCompleted == false)
            revert();

        Voter storage voter = voters[msg.sender];

        if(voter.hasRightToVote == false)
            revert();

        if(voter.voted == true)
            votingOptions[voter.votedFor].numberVotes -= 1;

        voter.voted = true;
        voter.votedFor = idx;

        votingOptions[idx].numberVotes += 1;

    }

    function getNumberRegisteredVoters() public view returns (uint)
    {
        return numberRegisteredVoters;
    }

    function getNameVotingOption(uint idx) public view returns (string memory)
    {
        return votingOptions[idx].name;
    }

    function getOptionsCompleted() public view returns (bool)
    {
        return optionsCompleted;
    }

    function getVotingEndTime() public view returns (uint)
    {
        return votingEndTime;
    }

    function getVotingTitle() public view returns (string memory)
    {
        return votingTitle;
    }

    function getVotingOptionsLength() public view returns (uint)
    {
        return votingOptions.length;
    }

    function getNumberVotesVotingOption(uint idx) public view returns (uint)
    {
        return votingOptions[idx].numberVotes;
    }
}