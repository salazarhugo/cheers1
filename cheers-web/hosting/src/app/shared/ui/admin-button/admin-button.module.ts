import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { AdminButtonRoutingModule } from './admin-button-routing.module';
import { AdminButtonComponent } from './admin-button.component';
import {MatButtonModule} from "@angular/material/button";


@NgModule({
    declarations: [
        AdminButtonComponent
    ],
    exports: [
        AdminButtonComponent
    ],
    imports: [
        CommonModule,
        AdminButtonRoutingModule,
        MatButtonModule
    ]
})
export class AdminButtonModule { }
