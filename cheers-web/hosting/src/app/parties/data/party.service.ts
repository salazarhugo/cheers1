import {Injectable} from '@angular/core';
import {ApiService} from "../../shared/data/services/api.service";
import {AngularFirestore} from "@angular/fire/compat/firestore";
import {firstValueFrom, lastValueFrom} from "rxjs";
import {Ticket} from "../../shared/data/models/ticket.model";
import {Party} from "../../shared/data/models/party.model";

@Injectable({
    providedIn: 'root'
})
export class PartyService {

    constructor(
        private api: ApiService,
        private fs: AngularFirestore,
    ) {
    }

    deleteParty(id: string) {
        return this.api.deleteParty(id)
    }

    createParty(party: Party) {
        return this.api.createParty(party)
    }

    getPartyTickets(id: string) {
        return firstValueFrom(this.fs.collection("ticketing").doc(id).collection<Ticket>("tickets").valueChanges())
    }

    interested(id: string) {
        return this.api.interestParty(id)
    }

    uninterested(id: string) {
        return this.api.uninterestParty(id)
    }

    getParty(id: string) {
        return this.api.getParty(id)
    }

    getMyParties() {
        return this.api.getMyParties()
    }

    getPartyFeed() {
        return this.api.getPartyFeed()
    }
}
