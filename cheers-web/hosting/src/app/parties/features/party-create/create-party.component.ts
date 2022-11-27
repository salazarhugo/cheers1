import {AfterViewInit, Component, ElementRef, OnInit, ViewChild} from '@angular/core';
import {UntypedFormBuilder, FormControl, FormGroup, Validator, Validators} from "@angular/forms";
import {debounceTime, distinctUntilChanged, finalize, fromEvent, lastValueFrom, map, Observable, of} from "rxjs";
import {User} from "../../../shared/data/models/user.model";
import {UserService} from "../../../shared/data/services/user.service";
import {MatSnackBar} from "@angular/material/snack-bar";
import {PartyService} from "../../data/party.service";
import {Ticket} from "../../../shared/data/models/ticket.model";
import {AngularFirestore} from "@angular/fire/compat/firestore";
import {ActivatedRoute, Router} from "@angular/router";
import {AngularFireStorage} from "@angular/fire/compat/storage";
import {Feature, MapboxGeocodingService} from "../../../shared/data/services/mapbox-geocoding.service";
import {Privacy} from "../../../shared/data/enum/privacy.enum";
import {Timestamp} from "google-protobuf/google/protobuf/timestamp_pb";
import {Party} from "../../../shared/data/models/party.model";
import {ActivityItemRoutingModule} from "../../../activities/ui/activity-item/activity-item-routing.module";

@Component({
    selector: 'app-create-party-detail',
    templateUrl: './create-party.component.html',
    styleUrls: ['./create-party.component.sass']
})
export class CreatePartyComponent implements OnInit {

    isLoading: boolean = false

    constructor(
        private userService: UserService,
        private partyService: PartyService,
        private _snackBar: MatSnackBar,
        private fb: UntypedFormBuilder,
        private fs: AngularFirestore,
        private storage: AngularFireStorage,
        private router: Router,
        private route: ActivatedRoute,
        private mapboxService: MapboxGeocodingService,
    ) {
    }

    partyForm = new FormGroup({
        name: new FormControl('', Validators.required),
        description: new FormControl(''),
        startDate: new FormControl(new Date(), Validators.required),
        startTime: new FormControl(''),
        endDate: new FormControl(new Date()),
        endTime: new FormControl(''),
        bannerUrl: new FormControl(''),
        privacy: new FormControl(Privacy.FRIENDS),
        latitude: new FormControl(0),
        longitude: new FormControl(0),
        locationName: new FormControl(''),
    });


    ngOnInit(): void {
    }

    openSnackBar(message: string, action: string) {
        this._snackBar.open(message, action, {duration: 3000});
    }

    async onSubmit(partyForm: any) {
        this.isLoading = true

        let party: Party = Object.assign(new Party(), partyForm);

        try {
            console.log("HUGO")
            console.log(party)
            const response = await lastValueFrom(this.partyService.createParty(party))
            this.openSnackBar("Party created", 'OK')
            await this.router.navigate(['/manage/parties', response.party?.id!, 'tickets'])
        } catch (e) {
            console.log(e)
            this.openSnackBar("Failed to create party-detail", 'OK')
        }
        this.isLoading = false
    }
}
