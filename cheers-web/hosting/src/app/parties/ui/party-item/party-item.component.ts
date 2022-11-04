import {Component, Input, OnInit} from '@angular/core';
import {Party} from "../../../shared/data/models/party.model";
import {PartyService} from "../../data/party.service";

@Component({
    selector: 'app-party-item',
    templateUrl: './party-item.component.html',
    styleUrls: ['./party-item.component.sass']
})
export class PartyItemComponent implements OnInit {

    @Input() party!: Party

    constructor(
        private partyService: PartyService,
    ) {
    }

    ngOnInit(): void {
    }

    onDeleteParty() {
        this.partyService.deleteParty(this.party.id).subscribe(res => console.log(res))
    }

    onImgError(event: any) {
        event.target.src = 'assets/default_profile_picture.png';
    }

    cancelParty(id: string) {
        this.partyService.deleteParty(id).subscribe(res => console.log(res))
    }
}
