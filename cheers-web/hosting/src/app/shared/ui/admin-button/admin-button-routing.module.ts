import { NgModule } from '@angular/core';
import { RouterModule, Routes } from '@angular/router';
import {AdminButtonComponent} from "./admin-button.component";

const routes: Routes = [{path: '', component: AdminButtonComponent}];

@NgModule({
  imports: [RouterModule.forChild(routes)],
  exports: [RouterModule]
})
export class AdminButtonRoutingModule { }
