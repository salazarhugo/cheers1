import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';

import {EditRoutingModule} from './edit-routing.module';
import {EditComponent} from "./edit.component";
import {FlexModule} from "@angular/flex-layout";
import {MatLegacyFormFieldModule as MatFormFieldModule} from "@angular/material/legacy-form-field";
import {MatLegacyProgressSpinnerModule as MatProgressSpinnerModule} from "@angular/material/legacy-progress-spinner";
import {ReactiveFormsModule} from "@angular/forms";
import {MatLegacyInputModule as MatInputModule} from "@angular/material/legacy-input";
import {MatLegacyButtonModule as MatButtonModule} from "@angular/material/legacy-button";


@NgModule({
    declarations: [EditComponent],
    imports: [
        CommonModule,
        EditRoutingModule,
        FlexModule,
        MatFormFieldModule,
        ReactiveFormsModule,
        MatProgressSpinnerModule,
        MatInputModule,
        MatButtonModule
    ],
    exports: [EditComponent],
})
export class EditModule {
}
