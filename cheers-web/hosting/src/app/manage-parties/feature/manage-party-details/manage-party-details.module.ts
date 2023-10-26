import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { ManagePartyDetailsRoutingModule } from './manage-party-details-routing.module';
import { ManagePartyDetailsComponent } from './manage-party-details.component';
import {FlexModule} from "@angular/flex-layout";
import {ReactiveFormsModule} from "@angular/forms";
import {NgxDropzoneModule} from "ngx-dropzone";
import {MatFormFieldModule} from "@angular/material/form-field";
import {MaterialModule} from "../../../material/material.module";


@NgModule({
  declarations: [
    ManagePartyDetailsComponent
  ],
    imports: [
        CommonModule,
        ManagePartyDetailsRoutingModule,
        FlexModule,
        MatFormFieldModule,
        ReactiveFormsModule,
        NgxDropzoneModule,
        MaterialModule
    ]
})
export class ManagePartyDetailsModule { }
