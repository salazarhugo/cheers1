import {Component, OnInit} from '@angular/core';
import {finalize, firstValueFrom, Observable, of} from "rxjs";
import {User} from "../../../shared/data/models/user.model";
import {UserService} from "../../../shared/data/services/user.service";
import {UntypedFormControl, UntypedFormGroup} from "@angular/forms";
import {AngularFireAuth} from "@angular/fire/compat/auth";
import {AngularFireStorage} from "@angular/fire/compat/storage";
import {MatSnackBar} from "@angular/material/snack-bar";


@Component({
    selector: 'app-user-edit',
    templateUrl: './edit.component.html',
    styleUrls: ['./edit.component.sass']
})
export class EditComponent implements OnInit {

    $user: Observable<User | null> = of(null)
    isLoading = false
    profileForm = new UntypedFormGroup({
        picture: new UntypedFormControl(''),
        name: new UntypedFormControl(''),
        username: new UntypedFormControl(''),
        email: new UntypedFormControl(''),
        website: new UntypedFormControl(''),
        bio: new UntypedFormControl(''),
    });

    constructor(
        private userService: UserService,
        private afAuth: AngularFireAuth,
        private _snackBar: MatSnackBar,
        private storage: AngularFireStorage,
    ) {
        this.$user = this.userService.user$
        this.$user.subscribe(user => {
            if (user != null)
                this.profileForm.patchValue(user)
        })
    }

    ngOnInit(): void {
    }

    async uploadFile(file: File): Promise<string> {
        const authUser = await firstValueFrom(this.afAuth.authState)
        const filePath = `${authUser!.uid}/${file.name}`;
        const storageRef = this.storage.ref(filePath);
        let res = (await this.storage.upload(filePath, file));
        return await res.ref.getDownloadURL()
    }

    async uploadPicture(file: File) {
        const url = await this.uploadFile(file)
        this.profileForm.get("picture")?.setValue(url)
        await this.userService.updateUser(this.profileForm.getRawValue()).toPromise()
    }

    addAttachment(fileInput: any) {
        const picture = fileInput.target.files[0] as File
        const reader = new FileReader();
        reader.readAsDataURL(picture);
        this.uploadPicture(picture)
    }

    openSnackBar(message: string, action: string) {
        this._snackBar.open(message, action, {duration: 3000});
    }

    async onSubmit() {
        this.isLoading = true
        try {
            await this.userService.updateUser(this.profileForm.getRawValue()).toPromise()
            this.openSnackBar("Profile saved", 'OK')
        } catch (e) {
            this.openSnackBar("Failed to save profile", 'OK')
        }
        this.isLoading = false
    }

    openFileChooser() {
        document.getElementById('upload-file')?.click()
    }
}
