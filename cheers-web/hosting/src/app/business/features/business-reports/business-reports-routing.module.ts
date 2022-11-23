import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {BusinessReportsComponent} from "./business-reports.component";

const routes: Routes = [{path: '', component: BusinessReportsComponent}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class BusinessReportsRoutingModule { }
