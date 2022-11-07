import {Component, OnInit} from '@angular/core';
import {ActivatedRoute, ParamMap, Router} from "@angular/router";
import {Party} from "../../../shared/data/models/party.model";
import {PartyService} from "../../data/party.service";
import {Clipboard} from '@angular/cdk/clipboard';
import {MatSnackBar} from "@angular/material/snack-bar";
import {PartyItem} from "../../../../gen/ts/cheers/party/v1/party_service";

@Component({
    selector: 'app-party-detail',
    templateUrl: './party.component.html',
    styleUrls: ['./party.component.sass']
})
export class PartyComponent implements OnInit {

    partyItem: PartyItem | undefined
    partyId: string | null

    constructor(
        public router: Router,
        private route: ActivatedRoute,
        private partyService: PartyService,
        private clipboard: Clipboard,
        private snackBar: MatSnackBar,
    ) {
    }

    ngOnInit(): void {
        this.route.paramMap.subscribe(async (params: ParamMap) => {
            const partyId = params.get("id")
            this.partyId = params.get("id")
            if (partyId)
                this.partyItem = await this.partyService.getParty(partyId).toPromise()
        })
    }

    onInterestedClick() {
        if (this.partyItem!.isInterested)
            this.partyService.uninterested(this.partyId!).subscribe(res => console.log(res))
        else
            this.partyService.interested(this.partyId!).subscribe(res => console.log(res))

        this.partyItem!.isInterested = !this.partyItem!.isInterested
    }

    copyLink() {
        this.clipboard.copy(window.location.href);
        this.snackBar.open("Link copied to clipboard", "Hide", {
            duration: 3000
        })
    }
}
