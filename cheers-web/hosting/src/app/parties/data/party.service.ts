import {Injectable} from '@angular/core';
import {ApiService} from "../../shared/data/services/api.service";
import {AngularFirestore} from "@angular/fire/compat/firestore";
import {BehaviorSubject, firstValueFrom, map, Observable} from "rxjs";
import {Ticket} from "../../shared/data/models/ticket.model";
import {
    AnswerPartyResponse,
    CreatePartyResponse,
    FeedPartyResponse,
    GetPartyItemResponse, ListGoingRequest, ListGoingResponse, ListPartyResponse,
    PartyItem, UpdatePartyRequest, UpdatePartyResponse, WatchStatus
} from "../../../gen/ts/cheers/party/v1/party_service";
import {HttpClient} from "@angular/common/http";
import {environment} from "../../../environments/environment";
import {Party, toParty} from "../../shared/data/models/party.model";
import {ListPartyOrdersResponse, Order as OrderGen} from "../../../gen/ts/cheers/order/v1/order_service";
import {UserItem} from "../../../gen/ts/cheers/type/user/user";

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

    transferParty(userId:string, partyId: string): Promise<void | Object> {
        return firstValueFrom(this.http.post(`${environment.GATEWAY_URL}/v1/parties/${partyId}/transfer`, {
            user_id: userId
        }))
    }

    listGoing(partyId: string): Promise<UserItem[]> {
        return firstValueFrom(this.http.get<ListGoingResponse>(`${environment.GATEWAY_URL}/v1/parties/${partyId}/going`)
            .pipe(map(r => r.users)))
    }

    answerParty(id: string, answer: WatchStatus): Promise<AnswerPartyResponse> {
        return firstValueFrom(this.http.post<AnswerPartyResponse>(`${environment.GATEWAY_URL}/v1/parties/${id}/answer`, {
            answer: answer
        }))
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
                privacy: party.privacy,
                bannerUrl: party.bannerUrl,
                locationName: party.locationName,
                start_date: party.startDate,
                end_date: party.startDate,
                latitude: party.latitude,
                longitude: party.longitude,
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

    getPartyOrders(partyId: string): Observable<OrderGen[]> {
        return this.http.get<ListPartyOrdersResponse>(`${environment.GATEWAY_URL}/v1/orders/party?party_id=${partyId}`)
            .pipe(map(p => p.orders));
    }

    getParty(id: string): Observable<Party> {
        return this.http.get<GetPartyItemResponse>(`${environment.GATEWAY_URL}/v1/parties/${id}/item`)
            .pipe(map(p => toParty(p.item!)));
    }

    getMyParties(userId: string): Observable<Party[]> {
        return this.http.get<ListPartyResponse>(`${environment.GATEWAY_URL}/v1/parties/${userId}?pageSize=10&page=1`)
            .pipe(map(res => res.items.map(p => toParty(p))));
    }

    getFeed(): Observable<Party[]> {
        return this.partyFeed$
    }

    getPartyFeed(): Observable<Party[]> {
        return this.http.get<FeedPartyResponse>(`${environment.GATEWAY_URL}/v1/parties/feed?pageSize=10&page=0`)
            .pipe(map(res => res.items.map(p => toParty(p))));
    }
}
