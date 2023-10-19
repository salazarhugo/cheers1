import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CheersButtonComponent } from './cheers-button.component';

describe('CheersButtonComponent', () => {
  let component: CheersButtonComponent;
  let fixture: ComponentFixture<CheersButtonComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CheersButtonComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CheersButtonComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
