import { ComponentFixture, TestBed } from '@angular/core/testing';

import { ManagePartyBasicinfoComponent } from './manage-party-basicinfo.component';

describe('ManagePartyBasicinfoComponent', () => {
  let component: ManagePartyBasicinfoComponent;
  let fixture: ComponentFixture<ManagePartyBasicinfoComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ ManagePartyBasicinfoComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(ManagePartyBasicinfoComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
