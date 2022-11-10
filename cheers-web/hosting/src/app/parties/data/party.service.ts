import {Injectable} from '@angular/core';
import {ApiService} from "../../shared/data/services/api.service";
import {AngularFirestore} from "@angular/fire/compat/firestore";
import {BehaviorSubject, firstValueFrom, lastValueFrom, map, Observable, of} from "rxjs";
import {Ticket} from "../../shared/data/models/ticket.model";
import {
    CreatePartyResponse,
    FeedPartyResponse,
    GetPartyItemResponse,
    PartyItem, UpdatePartyRequest, UpdatePartyResponse
} from "../../../gen/ts/cheers/party/v1/party_service";
import {HttpClient} from "@angular/common/http";
import {environment} from "../../../environments/environment";
import {Party, toParty} from "../../shared/data/models/party.model";
import {Order, toOrder} from "../../shared/data/models/order.model";
import {ListOrderResponse} from "../../../gen/ts/cheers/order/v1/order_service";

@Injectable({
    providedIn: 'root'
})
export class PartyService {

    private partyFeed$ = new BehaviorSubject<Party[]>([])
    private party$ = new BehaviorSubject<Party>(new Party())

    constructor(
        private api: ApiService,
        private fs: AngularFirestore,
        private http: HttpClient,
    ) {
    }

    setParty(party: Party) {
        this.party$.next(party)
    }

    getManagedParty() {
        return this.party$
    }

    deleteParty(id: string) {
        return this.api.deleteParty(id)
    }

    updateParty(party: any): Observable<UpdatePartyResponse> {
        return this.http.patch<UpdatePartyResponse>(`${environment.GATEWAY_URL}/v1/parties`, {
            party: party
        })
    }

    createParty(party: Party): Observable<CreatePartyResponse> {
        return this.http.post<CreatePartyResponse>(`${environment.GATEWAY_URL}/v1/parties`, {
            party: {
                name: party.name,
                description: party.description,
                address: party.address,
                privacy: "FRIENDS",
                bannerUrl: party.bannerUrl,
                locationName: party.locationName,
            }
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

    getPartyOrders(partyId: string): Observable<Order[]> {
        return this.http.get<ListOrderResponse>(`${environment.GATEWAY_URL}/v1/orders/list?party_id=${partyId}`)
            .pipe(map(p => p.orders.map(o => toOrder(o))));
    }

    getParty(id: string): Observable<Party> {
        return this.http.get<GetPartyItemResponse>(`${environment.GATEWAY_URL}/v1/parties/${id}/item`)
            .pipe(map(p => toParty(p.item!)));
    }

    getMyParties(): Observable<Party[]> {
        return this.http.get<PartyItem[]>(`${environment.GATEWAY_URL}/v1/parties/my?pageSize=10&page=0`)
            .pipe(map(res => res.map(p => toParty(p))));
    }

    getFeed(): Observable<Party[]> {
        return this.partyFeed$
    }

    getPartyFeed(): Observable<Party[]> {
        return this.http.get<FeedPartyResponse>(`${environment.GATEWAY_URL}/v1/parties/feed?pageSize=10&page=0`)
            .pipe(map(res => res.items.map(p => toParty(p))));
    }
}
