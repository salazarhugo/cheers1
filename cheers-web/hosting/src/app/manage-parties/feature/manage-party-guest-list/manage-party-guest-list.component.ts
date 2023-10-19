import {Component, OnInit} from '@angular/core';
import {PartyService} from "../../../parties/data/party.service";
import {firstValueFrom} from "rxjs";
import {Party} from "../../../shared/data/models/party.model";
import {User} from "../../../shared/data/models/user.model";

@Component({
    selector: 'app-manage-party-guest-list',
    templateUrl: './manage-party-guest-list.component.html',
    styleUrls: ['./manage-party-guest-list.component.sass']
})
export class ManagePartyGuestListComponent implements OnInit {

    users: User[] = []
    party: Party

    constructor(
        private partyService: PartyService,
    ) {
    }

    async ngOnInit() {
        const party = await firstValueFrom(this.partyService.getManagedParty())
        this.party = party
        this.users = await this.partyService.listGoing(party.id)
    }

}
