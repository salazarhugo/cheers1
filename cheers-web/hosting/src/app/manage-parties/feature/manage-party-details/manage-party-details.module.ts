import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ManagePartyDetailsRoutingModule } from './manage-party-details-routing.module';
import { ManagePartyDetailsComponent } from './manage-party-details.component';
import {FlexModule} from "@angular/flex-layout";
import {MatButtonModule} from "@angular/material/button";
import {MatFormFieldModule} from "@angular/material/form-field";
import {ReactiveFormsModule} from "@angular/forms";
import {MatInputModule} from "@angular/material/input";
import {NgxDropzoneModule} from "ngx-dropzone";
import {MatDividerModule} from "@angular/material/divider";
import {MatProgressSpinnerModule} from "@angular/material/progress-spinner";


@NgModule({
  declarations: [
    ManagePartyDetailsComponent
  ],
    imports: [
        CommonModule,
        ManagePartyDetailsRoutingModule,
        FlexModule,
        MatButtonModule,
        MatFormFieldModule,
        ReactiveFormsModule,
        MatInputModule,
        NgxDropzoneModule,
        MatDividerModule,
        MatProgressSpinnerModule
    ]
})
export class ManagePartyDetailsModule { }
