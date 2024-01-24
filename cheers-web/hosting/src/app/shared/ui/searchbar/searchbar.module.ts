import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';

import { SearchbarRoutingModule } from './searchbar-routing.module';
import { SearchbarComponent } from './searchbar.component';
import {MatFormFieldModule} from "@angular/material/form-field";
import {MatIconModule} from "@angular/material/icon";
import {FormsModule} from "@angular/forms";
import {MatInputModule} from "@angular/material/input";


@NgModule({
  declarations: [
    SearchbarComponent
  ],
  exports: [
    SearchbarComponent
  ],
  imports: [
    CommonModule,
    SearchbarRoutingModule,
    MatFormFieldModule,
    MatIconModule,
    FormsModule,
    MatInputModule
  ]
})
export class SearchbarModule { }
