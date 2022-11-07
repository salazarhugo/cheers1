import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, ParamMap} from "@angular/router";
import {PartyService} from "../../../parties/data/party.service";
import {Party} from "../../../shared/data/models/party.model";

@Component({
    selector: 'app-manage-party-sidenav',
    templateUrl: './manage-party-sidenav.component.html',
    styleUrls: ['./manage-party-sidenav.component.sass']
})
export class ManagePartySidenavComponent implements OnInit {

    party: Party | undefined

    items = [
        {
            title: "Basic Info",
            icon: "check_circle",
            routerLink: "basicinfo",
        },
        {
            title: "Details",
            icon: "check_circle",
            routerLink: "details",
        },
        {
            title: "Tickets",
            icon: "check_circle",
            routerLink: "tickets",
        },
        {
            title: "Publish",
            icon: "check_circle",
            routerLink: "publish",
        },
    ]

    constructor(
        private route: ActivatedRoute,
        private partyService: PartyService,
    ) {
        this.partyService.getManagedParty().subscribe(party => this.party = party)
    }

    ngOnInit(): void {
    }

}
