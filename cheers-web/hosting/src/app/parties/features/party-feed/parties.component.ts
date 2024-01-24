import {Component, OnInit} from '@angular/core';
import {concat, firstValueFrom, map, merge, Observable, of, Subject} from "rxjs";
import {UserService} from "../../../shared/data/services/user.service";
import {PartyService} from "../../data/party.service";
import {Party} from "../../../shared/data/models/party.model";
import {UserModel} from "../../../shared/data/models/user.model";
import {KeyValue} from "@angular/common";

@Component({
    selector: 'app-parties',
    templateUrl: './parties.component.html',
    styleUrls: ['./parties.component.sass']
})
export class PartiesComponent implements OnInit {

    $user: Observable<UserModel | null> = of(null)
    partyList: Party[] = []
    parties: Map<string, Party[]>
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
            this.parties = this.groupPartiesByDay(this.partyList)
            console.log(this.parties)
        })
    }

    // Preserve original property order
    originalOrder = (a: KeyValue<string, Party[]>, b: KeyValue<string, Party[]>): number => {
        return 0;
    }

    // Function to group parties by startDate day
    groupPartiesByDay(parties: Party[]): Map<string, Party[]> {
        const groupedParties = parties.reduce((groups, party) => {
            // Convert startDate (epoch) to a JavaScript Date object
            const startDate = new Date(party.startDate * 1000);

            // Get the day in a human-readable format, e.g., "Today", "Tomorrow", or a date string
            let day: string;
            const today = new Date();
            const tomorrow = new Date(today);
            tomorrow.setDate(today.getDate() + 1);

            if (startDate.toDateString() === today.toDateString()) {
                day = 'Today';
            } else if (startDate.toDateString() === tomorrow.toDateString()) {
                day = 'Tomorrow';
            } else {
                // Use the full date if not today or tomorrow
                day = startDate.toDateString();
            }

            // Initialize the group array if it doesn't exist
            if (!groups.has(day)) {
                groups.set(day, []);
            }

            // Add the party to the corresponding day's group
            groups.get(day)?.push(party);

            return groups;
        }, new Map<string, Party[]>());

        return groupedParties;
    }
}
