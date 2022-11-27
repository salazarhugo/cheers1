import {Component, ElementRef, EventEmitter, Input, OnInit, Output, ViewChild} from '@angular/core';
import {FormGroup, UntypedFormBuilder} from "@angular/forms";
import {Ticket} from "../../../shared/data/models/ticket.model";
import {Feature, MapboxGeocodingService} from "../../../shared/data/services/mapbox-geocoding.service";
import {debounceTime, distinctUntilChanged, finalize, fromEvent, map, Observable, of} from "rxjs";
import {UserService} from "../../../shared/data/services/user.service";
import {PartyService} from "../../data/party.service";
import {MatSnackBar} from "@angular/material/snack-bar";
import {AngularFirestore} from "@angular/fire/compat/firestore";
import {AngularFireStorage} from "@angular/fire/compat/storage";
import {Router} from "@angular/router";
import {User} from "../../../shared/data/models/user.model";
import {Privacy} from "../../../shared/data/enum/privacy.enum";
import {Party} from "../../../shared/data/models/party.model";

@Component({
    selector: 'app-party-form',
    templateUrl: './party-form.component.html',
    styleUrls: ['./party-form.component.sass', '../../features/party-detail/party.component.sass']
})
export class PartyFormComponent implements OnInit {
    @Input() partyForm!: FormGroup;
    @Output() onSubmit = new EventEmitter<Party>();

    privacyList = Object.values(Privacy)

    isLoading = false
    hasEndDate = false
    tickets: Ticket[] = []
    url: string = ""
    selectedLocation: Feature | undefined
    locationResults: Observable<Feature[]>;
    @ViewChild('searchInput', {static: true}) searchInput!: ElementRef;
    $user: Observable<User | null> = of(null)


    constructor(
        private fb: UntypedFormBuilder,
        private userService: UserService,
        private partyService: PartyService,
        private _snackBar: MatSnackBar,
        private fs: AngularFirestore,
        private storage: AngularFireStorage,
        private route: Router,
        private mapboxService: MapboxGeocodingService,
    ) {
        this.$user = this.userService.user$
    }

    toUnixTimestamp(date: Date): number {
        return Math.floor(date.getTime() / 1000);
    }
    submit() {
        const party = this.partyForm.getRawValue()
        const hours = party.startTime.split(":")[0]
        const minutes = party.startTime.split(":")[1]
        party.startDate = this.toUnixTimestamp(party.startDate) + hours * 3600 + minutes * 60
        party.endDate = this.toUnixTimestamp(party.endDate)
        party.privacy = party.privacy.toUpperCase()

        this.onSubmit.emit(party);
        this.partyForm.reset();
    }

    ngOnInit(): void {
    }

    ngAfterViewInit(): void {
        fromEvent(this.searchInput.nativeElement, 'keyup').pipe(
            map((event: any) => event.target.value),
            debounceTime(500),
            distinctUntilChanged()
        ).subscribe((query: string) => {
            this.searchPlace(query)
        });
    }

    selectPlace(feature: Feature): void {
        this.partyForm.get('latitude')?.setValue(feature.geometry.coordinates[1])
        this.partyForm.get('longitude')?.setValue(feature.geometry.coordinates[0])
        this.partyForm.get('locationName')?.setValue(feature.place_name)
    }

    searchPlace(query: string): void {
        this.locationResults = this.mapboxService.searchPlace(query)
    }

    handleClick() {
        document.getElementById('upload-file')?.click()
    }

    getOptionText(option: any) {
        return option.place_name;
    }

    openSnackBar(message: string, action: string) {
        this._snackBar.open(message, action, {duration: 3000});
    }

}
