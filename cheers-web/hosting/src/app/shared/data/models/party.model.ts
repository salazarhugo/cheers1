import {PartyItem} from "../../../../gen/ts/cheers/party/v1/party_service";

export class Party {
    isHost: boolean = false
    hostName: string = ""
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
    mutualProfilePictureUrls: string[] = []
    id: string = ""
    goingCount: number = 0
    participants: string[] = []
    longitude: number = 0.0
    mutualUsernames: string[] = []
    address: string = ""
    locationName: string = ""
    going: boolean = false
    created: number = 0
    showOnMap: boolean = false
    hostId: string = ""
    interestedCount: number = 0
    accountId: string = ""
    name: string = ""
    interested: boolean = false
    startDate: Date = new Date()
}

export function toParty(value: PartyItem): Party {
    const party = new Party()
    Object.assign(party, value.party)
    party.goingCount = value.goingCount
    party.interestedCount = value.interestedCount
    party.goingCount = value.invitedCount
    party.isHost = value.isCreator
    return party
}