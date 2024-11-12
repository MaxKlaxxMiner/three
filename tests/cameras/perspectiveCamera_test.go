package cameras

import (
	"github.com/MaxKlaxxMiner/three"
	"testing"
)

func TestPerspectiveCameraInstancing(t *testing.T) {
	c := three.NewPerspectiveCameraDefault()
	if c.Fov != 50 {
		t.Errorf("NewPerspectiveCameraDefault().Fov failed: expected (%f), got (%f)", 50.0, c.Fov)
	}
	if c.Zoom != 1 {
		t.Errorf("NewPerspectiveCameraDefault().Zoom failed: expected (%f), got (%f)", 1.0, c.Zoom)
	}
	if c.Near != 0.1 {
		t.Errorf("NewPerspectiveCameraDefault().Near failed: expected (%f), got (%f)", 0.1, c.Near)
	}
	if c.Far != 2000 {
		t.Errorf("NewPerspectiveCameraDefault().Far failed: expected (%f), got (%f)", 2000.0, c.Far)
	}
	if c.Focus != 10 {
		t.Errorf("NewPerspectiveCameraDefault().Focus failed: expected (%f), got (%f)", 10.0, c.Focus)
	}
	if c.Aspect != 1 {
		t.Errorf("NewPerspectiveCameraDefault().Aspect failed: expected (%f), got (%f)", 1.0, c.Aspect)
	}
	//todo: c.View
	if c.FilmGauge != 35 {
		t.Errorf("NewPerspectiveCameraDefault().FilmGauge failed: expected (%f), got (%f)", 35.0, c.FilmGauge)
	}
	if c.FilmOffset != 0 {
		t.Errorf("NewPerspectiveCameraDefault().FilmOffset failed: expected (%f), got (%f)", 0.0, c.FilmOffset)
	}
	c2 := &three.PerspectiveCamera{}
	if c2.Fov != 0 || c2.Zoom != 0 || c2.Near != 0 || c2.Far != 0 || c2.Focus != 0 || c2.Aspect != 0 || c2.FilmGauge != 0 || c2.FilmOffset != 0 {
		t.Errorf("copy failed zero values")
	}
	c2.Copy(c, true)
	if c2.Fov != 50 || c2.Zoom != 1 || c2.Near != 0.1 || c2.Far != 2000 || c2.Focus != 10 || c2.Aspect != 1 || c2.FilmGauge != 35 || c2.FilmOffset != 0 {
		t.Errorf("copy failed no expected values")
	}
	c3 := three.NewPerspectiveCamera(75, 1.2, 0.5, 1000.0)
	if c3.Fov != 75 || c3.Zoom != 1 || c3.Near != 0.5 || c3.Far != 1000 || c3.Focus != 10 || c3.Aspect != 1.2 || c3.FilmGauge != 35 || c3.FilmOffset != 0 {
		t.Errorf("c3 instance failed no expected values")
	}
}

func TestPerspectiveCameraValues(t *testing.T) {
	c := three.NewPerspectiveCameraDefault()
	if c.Fov != 50 {
		t.Errorf("NewPerspectiveCameraDefault().Fov failed: expected (%f), got (%f)", 50.0, c.Fov)
	}
	if c.GetFocalLength() != 37.52887110891728 {
		t.Errorf("c.GetFocalLength() failed: expected (%f), got (%f)", 37.52887110891728, c.GetFocalLength())
	}
	c.SetFocalLength(80.0)
	if c.Fov != 24.67817455665239 {
		t.Errorf("c.SetFocalLength(%f) failed: expected Fov (%f), got (%f)", 80.0, 24.67817455665239, c.Fov)
	}
	if c.GetFocalLength() != 80.0 {
		t.Errorf("c.GetFocalLength() failed: expected (%f), got (%f)", 80.0, c.GetFocalLength())
	}
	if c.GetEffectiveFOV() != 24.67817455665239 {
		t.Errorf("c.GetEffectiveFOV() failed: expected (%f), got (%f)", 24.67817455665239, c.GetEffectiveFOV())
	}
	if c.GetFilmWidth() != 35 {
		t.Errorf("c.GetFilmWidth() failed: expected (%f), got (%f)", 35.0, c.GetFilmWidth())
	}
	if c.GetFilmHeight() != 35 {
		t.Errorf("c.GetFilmHeight() failed: expected (%f), got (%f)", 35.0, c.GetFilmHeight())
	}
}

//		QUnit.todo( 'setViewOffset', ( assert ) => {
//
//			assert.ok( false, 'everything\'s gonna be alright' );
//
//		} );
//
//		QUnit.todo( 'clearViewOffset', ( assert ) => {
//
//			assert.ok( false, 'everything\'s gonna be alright' );
//
//		} );
//
//		QUnit.test( 'updateProjectionMatrix', ( assert ) => {
//
//			const cam = new PerspectiveCamera( 75, 16 / 9, 0.1, 300.0 );
//
//			// updateProjectionMatrix is called in constructor
//			const m = cam.projectionMatrix;
//
//			// perspective projection is given my the 4x4 Matrix
//			// 2n/r-l		0			l+r/r-l				 0
//			//   0		2n/t-b	t+b/t-b				 0
//			//   0			0		-(f+n/f-n)	-(2fn/f-n)
//			//   0			0				-1					 0
//
//			// this matrix was calculated by hand via glMatrix.perspective(75, 16 / 9, 0.1, 300.0, pMatrix)
//			// to get a reference matrix from plain WebGL
//			const reference = new Matrix4().set(
//				0.7330642938613892, 0, 0, 0,
//				0, 1.3032253980636597, 0, 0,
//				0, 0, - 1.000666856765747, - 0.2000666856765747,
//				0, 0, - 1, 0
//			);
//
//			// assert.ok( reference.equals(m) );
//			assert.ok( matrixEquals4( reference, m, 0.000001 ) );
//
//		} );
//
//		QUnit.todo( 'toJSON', ( assert ) => {
//
//			assert.ok( false, 'everything\'s gonna be alright' );
//
//		} );
//
//		// OTHERS
//		// TODO: clone is a camera methods that relied to copy method
//		QUnit.test( 'clone', ( assert ) => {
//
//			const near = 1,
//				far = 3,
//				aspect = 16 / 9,
//				fov = 90;
//
//			const cam = new PerspectiveCamera( fov, aspect, near, far );
//
//			const clonedCam = cam.clone();
//
//			assert.ok( cam.fov === clonedCam.fov, 'fov is equal' );
//			assert.ok( cam.aspect === clonedCam.aspect, 'aspect is equal' );
//			assert.ok( cam.near === clonedCam.near, 'near is equal' );
//			assert.ok( cam.far === clonedCam.far, 'far is equal' );
//			assert.ok( cam.zoom === clonedCam.zoom, 'zoom is equal' );
//			assert.ok( cam.projectionMatrix.equals( clonedCam.projectionMatrix ), 'projectionMatrix is equal' );
//
//		} );
//
//	} );
//
//} );
//		// see e.g. math/Matrix4.js
//		const matrixEquals4 = function ( a, b, tolerance ) {
//
//			tolerance = tolerance || 0.0001;
//			if ( a.elements.length != b.elements.length ) {
//
//				return false;
//
//			}
//
//			for ( let i = 0, il = a.elements.length; i < il; i ++ ) {
//
//				const delta = a.elements[ i ] - b.elements[ i ];
//				if ( delta > tolerance ) {
//
//					return false;
//
//				}
//
//			}
//
//			return true;
//
//		};
