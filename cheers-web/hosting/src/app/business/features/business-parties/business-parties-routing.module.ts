import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {BusinessPartiesComponent} from "./business-parties.component";

const routes: Routes = [{path: '', component: BusinessPartiesComponent}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class BusinessPartiesRoutingModule { }
