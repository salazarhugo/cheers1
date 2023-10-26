import {NgModule} from '@angular/core';
import {CommonModule} from '@angular/common';
import {FlexLayoutModule} from "@angular/flex-layout";
import {FormsModule, ReactiveFormsModule} from "@angular/forms";
import {MatStepperModule} from "@angular/material/stepper";
import {MatIconModule} from "@angular/material/icon";
import {MatDividerModule} from "@angular/material/divider";
import {MatSidenavModule} from "@angular/material/sidenav";
import {MatNativeDateModule, MatRippleModule} from "@angular/material/core";
import {MatToolbarModule} from "@angular/material/toolbar";
import {MatDatepickerModule} from "@angular/material/datepicker";
import {MatButtonToggleModule} from "@angular/material/button-toggle";
import {MatInputModule} from "@angular/material/input";
import {MatSnackBarModule} from "@angular/material/snack-bar";
import {MatDialogModule} from "@angular/material/dialog";
import {MatMenuModule} from "@angular/material/menu";
import {MatButtonModule} from "@angular/material/button";
import {MatProgressSpinnerModule} from "@angular/material/progress-spinner";
import {MatCardModule} from "@angular/material/card";
import {MatCheckboxModule} from "@angular/material/checkbox";
import {MatTabsModule} from "@angular/material/tabs";
import {MatFormFieldModule} from "@angular/material/form-field";
import {MatSelectModule} from "@angular/material/select";
import {MatTooltipModule} from "@angular/material/tooltip";
import {MatAutocompleteModule} from "@angular/material/autocomplete";
import {MatListModule} from "@angular/material/list";
import {MatSlideToggleModule} from "@angular/material/slide-toggle";


const components = [
    MatInputModule,
    MatSnackBarModule,
    MatDialogModule,
    MatMenuModule,
    FlexLayoutModule,
    MatButtonModule,
    MatProgressSpinnerModule,
    MatCardModule,
    MatCheckboxModule,
    MatTabsModule,
    MatSidenavModule,
    MatNativeDateModule,
    MatIconModule,
    MatToolbarModule,
    MatRippleModule,
    MatDividerModule,
    MatFormFieldModule,
    FormsModule,
    ReactiveFormsModule,
    MatDatepickerModule,
    MatSelectModule,
    MatStepperModule,
    MatTooltipModule,
    MatAutocompleteModule,
    MatButtonToggleModule,
    MatListModule,
    MatSlideToggleModule,
]

@NgModule({
    imports: [
        CommonModule,
        ...components
    ],
    exports: [
        ...components
    ]
})
export class MaterialModule {
}
