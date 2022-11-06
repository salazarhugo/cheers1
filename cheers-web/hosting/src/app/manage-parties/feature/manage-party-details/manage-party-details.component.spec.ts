import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ManagePartyDetailsComponent } from './manage-party-details.component';

describe('ManagePartyDetailsComponent', () => {
  let component: ManagePartyDetailsComponent;
  let fixture: ComponentFixture<ManagePartyDetailsComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ManagePartyDetailsComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ManagePartyDetailsComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
