import {Component, OnInit} from '@angular/core';
import {FormControl, FormGroup, Validators} from "@angular/forms";
import {finalize, lastValueFrom, Observable} from "rxjs";
import {AngularFireStorage} from "@angular/fire/compat/storage";
import {PartyService} from "../../../parties/data/party.service";
import {MatSnackBar} from "@angular/material/snack-bar";
import {Party, toParty} from "../../../shared/data/models/party.model";

@Component({
    selector: 'app-manage-party-details',
    templateUrl: './manage-party-details.component.html',
    styleUrls: ['./manage-party-details.component.sass']
})
export class ManagePartyDetailsComponent implements OnInit {

    partyForm = new FormGroup({
        id: new FormControl(''),
        description: new FormControl(''),
        summary: new FormControl(''),
        bannerUrl: new FormControl(''),
    });
    isLoading = false

    constructor(
        private storage: AngularFireStorage,
        private partyService: PartyService,
        private _snackBar: MatSnackBar,
    ) {
    }

    canSave: boolean = false
    initialValue: any

    ngOnInit(): void {
        this.partyService.getManagedParty().subscribe(party => {
            this.initialValue = {
                id: party.id,
                description: party.description,
                bannerUrl: party.bannerUrl,
            }
            this.partyForm.patchValue(this.initialValue)
        })
        this.partyForm.valueChanges.subscribe(form => {
            console.log(this.initialValue)
            console.log(form)
            this.canSave = JSON.stringify(form) !== JSON.stringify(this.initialValue)
        })
    }

    banner: File | null;

    onSelect(event: any) {
        console.log(event);
        this.banner = event.addedFiles[0]
    }

    onRemove(event: any) {
        this.banner = null
    }

    async pushFileToStorage(file: File): Promise<string> {
        const filePath = `parties/${file.name}`;
        const storageRef = this.storage.ref(filePath);
        let res = (await this.storage.upload(filePath, file));
        let urlDownloadImage = await res.ref.getDownloadURL();
        console.log(urlDownloadImage)
        return urlDownloadImage
    }

    openSnackBar(message: string, action: string) {
        this._snackBar.open(message, action, {duration: 3000});
    }

    onDiscard() {
        this.partyForm.patchValue(this.initialValue)
    }

    async onSave() {
        this.isLoading = true
        let form = this.partyForm.getRawValue()
        try {
            if (this.banner) {
                const url = await this.pushFileToStorage(this.banner)
                form.bannerUrl = url
            }
            const response = await lastValueFrom(this.partyService.updateParty(form))
            this.openSnackBar("Party updated", 'Hide')
        } catch (e) {
            console.log(e)
            this.openSnackBar("Failed to update party details", 'Hide')
        }
        this.isLoading = false
    }
}
