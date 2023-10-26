import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { BusinessOrdersRoutingModule } from './business-orders-routing.module';
import { BusinessOrdersComponent } from './business-orders.component';
import {OrderItemModule} from "../../../manage-parties/ui/order-item/order-item.module";
import {MatLegacyTableModule as MatTableModule} from "@angular/material/legacy-table";
import {MatSortModule} from "@angular/material/sort";
import {MatLegacyChipsModule as MatChipsModule} from "@angular/material/legacy-chips";
import {MatLegacyButtonModule as MatButtonModule} from "@angular/material/legacy-button";
import {MatIconModule} from "@angular/material/icon";
import {MatLegacyMenuModule as MatMenuModule} from "@angular/material/legacy-menu";
import {RelativeTimeModule} from "../../../shared/data/pipes/relative-time/relative-time.module";
import {MatLegacyFormFieldModule as MatFormFieldModule} from "@angular/material/legacy-form-field";
import {FormsModule} from "@angular/forms";
import {MatLegacyInputModule as MatInputModule} from "@angular/material/legacy-input";
import {SearchbarModule} from "../../../shared/ui/searchbar/searchbar.module";


@NgModule({
  declarations: [
    BusinessOrdersComponent
  ],
    imports: [
        CommonModule,
        BusinessOrdersRoutingModule,
        OrderItemModule,
        MatTableModule,
        MatSortModule,
        MatChipsModule,
        MatButtonModule,
        MatIconModule,
        MatMenuModule,
        RelativeTimeModule,
        MatFormFieldModule,
        FormsModule,
        MatInputModule,
        SearchbarModule
    ]
})
export class BusinessOrdersModule { }
