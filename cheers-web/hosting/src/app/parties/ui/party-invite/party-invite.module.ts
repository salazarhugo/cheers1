import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {PartyInviteRoutingModule} from './party-invite-routing.module';
import {PartyInviteComponent} from './party-invite.component';
import {MAT_LEGACY_DIALOG_DATA as MAT_DIALOG_DATA, MatLegacyDialogModule as MatDialogModule, MatLegacyDialogRef as MatDialogRef} from "@angular/material/legacy-dialog";
import {MatLegacyButtonModule as MatButtonModule} from "@angular/material/legacy-button";
import {MatLegacyChipsModule as MatChipsModule} from '@angular/material/legacy-chips';
import {UserItemModule} from "../../../users/ui/user-item/user-item.module";
import {MatLegacyRadioModule as MatRadioModule} from "@angular/material/legacy-radio";
import {MatLegacyCheckboxModule as MatCheckboxModule} from "@angular/material/legacy-checkbox";
import {MatLegacyProgressSpinnerModule as MatProgressSpinnerModule} from "@angular/material/legacy-progress-spinner";
import {FlexModule} from "@angular/flex-layout";

@NgModule({
    declarations: [
        PartyInviteComponent
    ],
    imports: [
        CommonModule,
        PartyInviteRoutingModule,
        MatDialogModule,
        MatButtonModule,
        MatChipsModule,
        UserItemModule,
        MatRadioModule,
        MatCheckboxModule,
        MatProgressSpinnerModule,
        FlexModule,
    ],
    exports: [
        PartyInviteComponent
    ],
    providers: [
        { provide: MatDialogRef, useValue: {} },
        { provide: MAT_DIALOG_DATA, useValue: {} },
    ],
})
export class PartyInviteModule {
}
