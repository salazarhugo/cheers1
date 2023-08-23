import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { SignUpRoutingModule } from './sign-up-routing.module';
import { SignUpComponent } from './sign-up.component';
import {SharedModule} from "../shared/shared.module";
import {FlexModule} from "@angular/flex-layout";


@NgModule({
  declarations: [
    SignUpComponent
  ],
  imports: [
    CommonModule,
    SignUpRoutingModule,
    SharedModule,
    FlexModule
  ]
})
export class SignUpModule { }
