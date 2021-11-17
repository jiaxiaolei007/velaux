import React from 'react';
import Translation from '../Translation';

import './index.less';

type Props = {
  message?: string | React.ReactNode;
  style?: {};
  iconWidth?: string;
};

const Empty = (props: Props) => {
  return (
    <div className="empty" style={props.style}>
      <div>
        <svg
          style={{ width: props.iconWidth }}
          viewBox="0 0 1024 1024"
          version="1.1"
          xmlns="http://www.w3.org/2000/svg"
          p-id="1890"
        >
          <path
            d="M512 0C229.261369 0 0 229.261369 0 512c0 282.738631 229.261369 512 512 512 282.738631 0 512-229.261369 512-512C1024 229.261369 794.738631 0 512 0L512 0zM512 999.82009c-269.049475 0-487.948026-218.898551-487.948026-487.948026 0-269.049475 218.898551-487.948026 487.948026-487.948026 269.049475 0 487.948026 218.898551 487.948026 487.948026C999.82009 781.049475 781.049475 999.82009 512 999.82009L512 999.82009zM512 999.82009"
            p-id="1891"
            fill="#bfbfbf"
          ></path>
          <path
            d="M223.632184 430.888556c0 16.375812 6.78061 32.623688 18.294853 44.265867 11.642179 11.642179 27.890055 18.294853 44.265867 18.294853 16.375812 0 32.623688-6.78061 44.265867-18.294853 11.642179-11.642179 18.294853-27.890055 18.294853-44.265867 0-16.375812-6.78061-32.751624-18.294853-44.265867-11.642179-11.642179-27.890055-18.294853-44.265867-18.294853-16.375812 0-32.751624 6.78061-44.265867 18.294853C230.412794 398.264868 223.632184 414.512744 223.632184 430.888556L223.632184 430.888556zM223.632184 430.888556"
            p-id="1892"
            fill="#bfbfbf"
          ></path>
          <path
            d="M674.606697 424.61969c0 16.503748 6.78061 33.007496 18.550725 44.777611 11.642179 11.642179 28.145927 18.550725 44.649675 18.550725 16.503748 0 33.007496-6.78061 44.777611-18.550725 11.770115-11.770115 18.550725-28.145927 18.550725-44.777611 0-16.503748-6.78061-33.007496-18.550725-44.777611-11.642179-11.642179-28.145927-18.550725-44.777611-18.550725-16.503748 0-33.007496 6.78061-44.649675 18.550725C681.387306 391.612194 674.606697 408.115942 674.606697 424.61969L674.606697 424.61969zM674.606697 424.61969"
            p-id="1893"
            fill="#bfbfbf"
          ></path>
          <path
            d="M385.471264 275.702149c-4.605697-11.642179-17.3993-17.527236-28.785607-13.049475l-153.77911 60.257871c-11.386307 4.477761-16.75962 17.527236-12.281859 29.041479l4.093953 10.490755c4.605697 11.642179 17.3993 17.527236 28.785607 13.049475l153.77911-60.257871c11.386307-4.477761 16.75962-17.527236 12.281859-29.169415L385.471264 275.702149 385.471264 275.702149zM385.471264 275.702149"
            p-id="1894"
            fill="#bfbfbf"
          ></path>
          <path
            d="M830.688656 318.816592l-156.465767-52.965517c-11.514243-3.838081-24.17991 2.558721-28.145927 14.328836l-3.582209 10.746627c-3.966017 11.898051 2.046977 24.563718 13.561219 28.529735l156.465767 52.965517c11.514243 3.966017 24.051974-2.558721 28.145927-14.328836l3.582209-10.746627C848.343828 335.448276 842.202899 322.654673 830.688656 318.816592L830.688656 318.816592zM830.688656 318.816592"
            p-id="1895"
            fill="#bfbfbf"
          ></path>
          <path
            d="M696.995502 663.604198c-6.268866-9.083458-18.550725-11.514243-27.378311-5.501249l-120.515742 82.902549c-8.827586 6.14093-11.002499 18.422789-4.733633 27.634183l5.629185 8.187906c6.268866 9.083458 18.550725 11.514243 27.378311 5.501249l120.515742-82.902549c8.827586-6.14093 11.002499-18.422789 4.733633-27.634183L696.995502 663.604198 696.995502 663.604198zM696.995502 663.604198"
            p-id="1896"
            fill="#bfbfbf"
          ></path>
        </svg>
      </div>
      <div className="message">{props.message || <Translation>Empty Data</Translation>}</div>
    </div>
  );
};

export default Empty;
