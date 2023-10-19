import {PartyItem, WatchStatus as WatchStatusPb} from "../../../../gen/ts/cheers/party/v1/party_service";
import {User} from "./user.model";

export enum WatchStatus {
    GOING = "Going",
    INTERESTED = "Interested",
    NOT_INTERESTED = "Not interested",
    MAYBE = "Maybe",
}

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
    mutualGoing: User[]
    price: number = 0
    bannerUrl: string = ""
    id: string = ""
    goingCount: number = 0
    interestedCount: number = 0
    invitedCount: number = 0
    longitude: number = 0.0
    address: string = ""
    locationName: string = ""
    createTime: number = 0
    hostId: string = ""
    name: string = ""
    isNotInterested: boolean = true
    watchStatus: WatchStatus = WatchStatus.NOT_INTERESTED
}

export function toParty(value: PartyItem): Party {
    const party = new Party()
    Object.assign(party, value.party)
    party.watchStatus = toWatchStatus(value.viewerWatchStatus)
    party.isNotInterested = party.watchStatus == WatchStatus.NOT_INTERESTED
    party.goingCount = value.goingCount
    party.interestedCount = value.interestedCount
    party.invitedCount = value.invitedCount
    party.isHost = value.isCreator
    party.hostName = value.user?.name  || ""
    party.startDate = Number(value.party?.startDate)
    party.endDate = Number(value.party?.endDate)
    console.log(party)
    return party
}

function toWatchStatus(value: any): WatchStatus {
    switch (value) {
        case "GOING": return WatchStatus.GOING;
        case "UNWATCHED": return WatchStatus.NOT_INTERESTED;
        case "INTERESTED": return WatchStatus.INTERESTED;
        case "MAYBE": return WatchStatus.MAYBE;
        case "UNRECOGNIZED": return WatchStatus.NOT_INTERESTED;
        default: return WatchStatus.NOT_INTERESTED;
    }
}

export function toWatchStatusPb(value: WatchStatus): WatchStatusPb {
    switch (value) {
        case WatchStatus.INTERESTED: return WatchStatusPb.INTERESTED;
        case WatchStatus.MAYBE: return WatchStatusPb.MAYBE;
        case WatchStatus.GOING: return WatchStatusPb.GOING;
        default: return WatchStatusPb.UNWATCHED;
    }
}
