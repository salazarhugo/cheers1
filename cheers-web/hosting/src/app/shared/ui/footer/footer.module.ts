import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { FooterComponent } from './footer.component';
import {FlexModule} from "@angular/flex-layout";
import {MatLegacyButtonModule as MatButtonModule} from "@angular/material/legacy-button";
import {MatDividerModule} from "@angular/material/divider";



@NgModule({
  declarations: [
    FooterComponent
  ],
  exports: [
    FooterComponent
  ],
  imports: [
    CommonModule,
    FlexModule,
    MatButtonModule,
    MatDividerModule
  ]
})
export class FooterModule { }
