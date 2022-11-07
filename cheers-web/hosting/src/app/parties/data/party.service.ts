import {Injectable} from '@angular/core';
import {ApiService} from "../../shared/data/services/api.service";
import {AngularFirestore} from "@angular/fire/compat/firestore";
import {firstValueFrom, lastValueFrom, map, Observable} from "rxjs";
import {Ticket} from "../../shared/data/models/ticket.model";
import {CreatePartyRequest, FeedPartyResponse, PartyItem} from "../../../gen/ts/cheers/party/v1/party_service";
import {Party} from "../../../gen/ts/cheers/type/party/party";
import {HttpClient} from "@angular/common/http";
import {environment} from "../../../environments/environment";

@Injectable({
    providedIn: 'root'
})
export class PartyService {

    constructor(
        private api: ApiService,
        private fs: AngularFirestore,
        private http: HttpClient,
    ) {
    }

    deleteParty(id: string) {
        return this.api.deleteParty(id)
    }

    createParty(party: Party): Observable<Party> {
        return this.http.post<Party>(`${environment.GATEWAY_URL}/v1/parties`, {
            party: party
        })
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

    getParty(id: string): Observable<Party | null> {
        return this.api.getParty(id)
    }

    getPartyItem(id: string): Observable<PartyItem> {
        return this.api.getPartyItem(id)
    }


    getMyParties(): Observable<PartyItem[]> {
        return this.http.get<PartyItem[]>(`${environment.GATEWAY_URL}/v1/parties/my?pageSize=10&page=0`)
    }

    getPartyFeed(): Observable<Party[]> {
        return this.http.get<FeedPartyResponse>(`${environment.GATEWAY_URL}/v1/parties/feed?pageSize=10&page=0`)
            .pipe(map(res => res.items.map(p => {
                return new Par
            })))
    }
}
