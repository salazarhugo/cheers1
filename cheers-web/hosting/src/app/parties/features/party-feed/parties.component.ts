import {Component, OnInit} from '@angular/core';
import {concat, firstValueFrom, map, merge, Observable, of, Subject} from "rxjs";
import {UserService} from "../../../shared/data/services/user.service";
import {PartyService} from "../../data/party.service";
import {Party} from "../../../shared/data/models/party.model";
import {User} from "../../../shared/data/models/user.model";

@Component({
    selector: 'app-parties',
    templateUrl: './parties.component.html',
    styleUrls: ['./parties.component.sass']
})
export class PartiesComponent implements OnInit {

    $user: Observable<User | null> = of(null)
    partyList: Party[] = []
    private contactsList = [];
    $myParties: Observable<Party[] | null> = of(null)
    chips: string[] = ["All", "Mix", "Outside", "Indoor", "Gadgets", "Android", "Piano", "Cloud computing", "Nouveaute", "Echecs"];
    page = 1
    pageSize = 18
    endReached = false

    constructor(
        private userService: UserService,
        private partyService: PartyService,
    ) {
        this.$user = this.userService.user$
    }

    async ngOnInit() {
        this.loadParties()
    }

    onImgError(event: any) {
        event.target.src = 'assets/default_profile_picture.png';
    }

    loadParties() {
        this.partyService.getPartyFeed(this.page, this.pageSize).subscribe(parties => {
            this.endReached = parties.length < this.pageSize
            if (!this.endReached)
                this.page++;
            this.partyList.push(...parties)
        })
    }
}
