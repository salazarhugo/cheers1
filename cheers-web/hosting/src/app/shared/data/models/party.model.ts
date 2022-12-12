import {PartyAnswer, PartyItem} from "../../../../gen/ts/cheers/party/v1/party_service";

export class Party {
    isHost: boolean = false
    hostName: string = ""
    startDate: number = 0
    endDate: number = 0
    owner: boolean = false
    latitude: number = 0.0
    showGuestList: boolean = false
    description: string = ""
    privacy: string = ""
    type: string = ""
    mutualCount: number = 0
    price: number = 0
    bannerUrl: string = ""
    id: string = ""
    goingCount: number = 0
    interestedCount: number = 0
    invitedCount: number = 0
    longitude: number = 0.0
    address: string = ""
    locationName: string = ""
    going: boolean = false
    createTime: number = 0
    hostId: string = ""
    name: string = ""
    interested: boolean = false
}

export function toParty(value: PartyItem): Party {
    const party = new Party()
    Object.assign(party, value.party)
    party.going = value.answer.toString() == "GOING"
    party.interested = value.answer == PartyAnswer.INTERESTED
    party.goingCount = value.goingCount
    party.interestedCount = value.interestedCount
    party.invitedCount = value.invitedCount
    party.isHost = value.isCreator
    party.hostName = value.user?.name  || ""
    party.startDate = Number(value.party?.startDate)
    party.endDate = Number(value.party?.endDate)
    return party
}